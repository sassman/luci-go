lucicfg.config(config_dir = '.output')
lucicfg.config(tracked_files = ['*.cfg'])
lucicfg.config(fail_on_warnings = True)

luci.project(
    name = 'infra',

    buildbucket = 'cr-buildbucket.appspot.com',
    logdog = 'luci-logdog.appspot.com',
    milo = 'luci-milo.appspot.com',
    scheduler = 'luci-scheduler.appspot.com',
    swarming = 'chromium-swarm.appspot.com',

    acls = [
        acl.entry(
            roles = [
                acl.PROJECT_CONFIGS_READER,
                acl.LOGDOG_READER,
                acl.BUILDBUCKET_READER,
                acl.SCHEDULER_READER,
            ],
            groups = ['all'],
        ),
        acl.entry(
            roles = [
                acl.BUILDBUCKET_OWNER,
                acl.SCHEDULER_OWNER,
            ],
            groups = ['admins'],
        ),
    ],
)

luci.logdog(gs_bucket = 'chromium-luci-logdog')

luci.milo(
    logo = 'https://storage.googleapis.com/chrome-infra-public/logo/chrome-infra-logo-200x200.png',
    favicon = 'https://storage.googleapis.com/chrome-infra-public/logo/favicon.ico',
    monorail_project = 'tutu, all aboard',
    monorail_components = ['Stuff>Hard'],
    bug_summary = 'Bug summary',
    bug_description = 'Everything is broken',
)


# Recipes.

luci.recipe(
    name = 'main/recipe',
    cipd_package = 'recipe/bundles/main',
)


# CI bucket.

luci.bucket(
    name = 'ci',

    # Allow developers to force-launch CI builds through Scheduler, but not
    # directly through Buildbucket. The direct access to Buildbucket allows to
    # override almost all aspects of the builds (e.g. what recipe is used),
    # and Buildbucket totally ignores any concurrency limitations set in the
    # LUCI Scheduler configs. This makes direct Buildbucket access to CI buckets
    # dangerous. They usually have very small pool of machines, and these
    # machines are assumed to be running only "approved" code (being post-submit
    # builders).
    acls = [
        acl.entry(acl.SCHEDULER_TRIGGERER, groups = ['devs']),
    ],
)

luci.gitiles_poller(
    name = 'master-poller',
    bucket = 'ci',
    repo = 'https://noop.com',
    refs = [
        'refs/heads/master',
        'refs/tags/blah',
        'refs/branch-heads/\d+\.\d+',
    ],
    schedule = 'with 10s interval',
)

luci.builder(
    name = 'linux ci builder',
    bucket = 'ci',
    recipe = 'main/recipe',

    triggered_by = ['master-poller'],
    triggers = ['ci/generically named builder'],

    properties = {
        'prop1': 'val1',
        'prop2': ['val2', 123],
    },
    service_account = 'builder@example.com',

    caches = [
        swarming.cache('path1'),
        swarming.cache('path2', name='name2'),
        swarming.cache('path3', name='name3', wait_for_warm_cache=10*time.minute),
    ],
    execution_timeout = 3 * time.hour,

    dimensions = {
        'os': 'Linux',
        'builder': 'linux ci builder',  # no auto_builder_dimension
        'prefer_if_available': [
            swarming.dimension('first-choice', expiration=5*time.minute),
            swarming.dimension('fallback'),
        ],
    },
    priority = 80,
    swarming_tags = ['tag1:val1', 'tag2:val2'],
    expiration_timeout = time.hour,
    build_numbers = True,

    triggering_policy = scheduler.greedy_batching(
        max_concurrent_invocations=5,
        max_batch_size=10,
    )
)

luci.builder(
    name = 'generically named builder',
    bucket = 'ci',
    recipe = 'main/recipe',

    triggered_by = ['master-poller'],
)

luci.builder(
    name = 'cron builder',
    bucket = 'ci',
    recipe = 'main/recipe',
    schedule = '0 6 * * *',
)


# Try bucket.

luci.bucket(
    name = 'try',

    # Allow developers to launch try jobs directly with whatever parameters
    # they want. Try bucket is basically a free build farm for all developers.
    acls = [
        acl.entry(acl.BUILDBUCKET_TRIGGERER, groups='devs'),
    ],
)

luci.builder(
    name = 'linux try builder',
    bucket = 'try',
    recipe = 'main/recipe',
)

luci.builder(
    name = 'generically named builder',
    bucket = 'try',
    recipe = 'main/recipe',
)


# Inline definitions.


luci.builder(
    name = 'triggerer builder',
    bucket = luci.bucket(name = 'inline'),
    recipe = luci.recipe(
        name = 'inline/recipe',
        cipd_package = 'recipe/bundles/inline',
    ),

    service_account = 'builder@example.com',

    triggers = [
        luci.builder(
            name = 'triggered builder',
            bucket = 'inline',
            recipe = 'inline/recipe',
        ),
    ],

    triggered_by = [
        luci.gitiles_poller(
            name = 'inline poller',
            bucket = 'inline',
            repo = 'https://noop.com',
            refs = [
                'refs/heads/master',
                'refs/tags/blah',
                'refs/branch-heads/\d+\.\d+',
            ],
            schedule = 'with 10s interval',
        ),
    ],
)


# List views.


luci.list_view(
    name = 'List view',
    entries = [
        'cron builder',
        'ci/generically named builder',
        luci.list_view_entry(
            builder = 'linux ci builder',
            buildbot = 'master/builder',
        ),
    ],
)

luci.list_view_entry(
    list_view = 'List view',
    builder = 'inline/triggered builder',
)

luci.list_view_entry(
    list_view = 'List view',
    buildbot = 'master/very buildbot',
)


# Console views.


luci.console_view(
    name = 'Console view',
    title = 'CI Builders',
    header = {
        'links': [
            {'name': 'a', 'links': [{'text': 'a'}]},
            {'name': 'b', 'links': [{'text': 'b'}]},
        ],
    },
    repo = 'https://noop.com',
    refs = ['refs/tags/blah', 'refs/branch-heads/\d+\.\d+'],
    exclude_ref = 'refs/heads/master',
    include_experimental_builds = True,
    entries = [
        luci.console_view_entry(
            builder = 'linux ci builder',
            buildbot = 'master/builder',
            category = 'a|b',
            short_name = 'lnx',
        ),
        # An alias for luci.console_view_entry(**{...}).
        {'builder': 'cron builder', 'category': 'cron'},
    ],
)

luci.console_view_entry(
    console_view = 'Console view',
    builder = 'inline/triggered builder',
)

luci.console_view_entry(
    console_view = 'Console view',
    buildbot = 'master/very buildbot',
)


# CQ.

luci.cq(
    submit_max_burst = 10,
    submit_burst_delay = 10 * time.minute,
    draining_start_time = '2017-12-23T15:47:58Z',
    status_host = 'chromium-cq-status.appspot.com',
)


# Expect configs:
#
# === commit-queue.cfg
# draining_start_time: "2017-12-23T15:47:58Z"
# cq_status_host: "chromium-cq-status.appspot.com"
# submit_options: <
#   max_burst: 10
#   burst_delay: <
#     seconds: 600
#   >
# >
# ===
#
# === cr-buildbucket.cfg
# buckets: <
#   name: "ci"
#   acls: <
#     role: WRITER
#     group: "admins"
#   >
#   acls: <
#     group: "all"
#   >
#   swarming: <
#     builders: <
#       name: "cron builder"
#       swarming_host: "chromium-swarm.appspot.com"
#       recipe: <
#         name: "main/recipe"
#         cipd_package: "recipe/bundles/main"
#         cipd_version: "refs/heads/master"
#       >
#     >
#     builders: <
#       name: "generically named builder"
#       swarming_host: "chromium-swarm.appspot.com"
#       recipe: <
#         name: "main/recipe"
#         cipd_package: "recipe/bundles/main"
#         cipd_version: "refs/heads/master"
#       >
#     >
#     builders: <
#       name: "linux ci builder"
#       swarming_host: "chromium-swarm.appspot.com"
#       swarming_tags: "tag1:val1"
#       swarming_tags: "tag2:val2"
#       dimensions: "builder:linux ci builder"
#       dimensions: "os:Linux"
#       dimensions: "300:prefer_if_available:first-choice"
#       dimensions: "prefer_if_available:fallback"
#       recipe: <
#         name: "main/recipe"
#         cipd_package: "recipe/bundles/main"
#         cipd_version: "refs/heads/master"
#         properties_j: "prop1:\"val1\""
#         properties_j: "prop2:[\"val2\",123]"
#       >
#       priority: 80
#       execution_timeout_secs: 10800
#       expiration_secs: 3600
#       caches: <
#         name: "name2"
#         path: "path2"
#       >
#       caches: <
#         name: "name3"
#         path: "path3"
#         wait_for_warm_cache_secs: 600
#       >
#       caches: <
#         name: "path1"
#         path: "path1"
#       >
#       build_numbers: YES
#       service_account: "builder@example.com"
#     >
#   >
# >
# buckets: <
#   name: "inline"
#   acls: <
#     role: WRITER
#     group: "admins"
#   >
#   acls: <
#     group: "all"
#   >
#   swarming: <
#     builders: <
#       name: "triggered builder"
#       swarming_host: "chromium-swarm.appspot.com"
#       recipe: <
#         name: "inline/recipe"
#         cipd_package: "recipe/bundles/inline"
#         cipd_version: "refs/heads/master"
#       >
#     >
#     builders: <
#       name: "triggerer builder"
#       swarming_host: "chromium-swarm.appspot.com"
#       recipe: <
#         name: "inline/recipe"
#         cipd_package: "recipe/bundles/inline"
#         cipd_version: "refs/heads/master"
#       >
#       service_account: "builder@example.com"
#     >
#   >
# >
# buckets: <
#   name: "try"
#   acls: <
#     role: WRITER
#     group: "admins"
#   >
#   acls: <
#     group: "all"
#   >
#   acls: <
#     role: SCHEDULER
#     group: "devs"
#   >
#   swarming: <
#     builders: <
#       name: "generically named builder"
#       swarming_host: "chromium-swarm.appspot.com"
#       recipe: <
#         name: "main/recipe"
#         cipd_package: "recipe/bundles/main"
#         cipd_version: "refs/heads/master"
#       >
#     >
#     builders: <
#       name: "linux try builder"
#       swarming_host: "chromium-swarm.appspot.com"
#       recipe: <
#         name: "main/recipe"
#         cipd_package: "recipe/bundles/main"
#         cipd_version: "refs/heads/master"
#       >
#     >
#   >
# >
# ===
#
# === luci-logdog.cfg
# reader_auth_groups: "all"
# archive_gs_bucket: "chromium-luci-logdog"
# ===
#
# === luci-milo.cfg
# consoles: <
#   id: "List view"
#   name: "List view"
#   builders: <
#     name: "buildbucket/luci.infra.ci/cron builder"
#   >
#   builders: <
#     name: "buildbucket/luci.infra.ci/generically named builder"
#   >
#   builders: <
#     name: "buildbot/master/builder"
#     name: "buildbucket/luci.infra.ci/linux ci builder"
#   >
#   builders: <
#     name: "buildbucket/luci.infra.inline/triggered builder"
#   >
#   builders: <
#     name: "buildbot/master/very buildbot"
#   >
#   favicon_url: "https://storage.googleapis.com/chrome-infra-public/logo/favicon.ico"
#   builder_view_only: true
# >
# consoles: <
#   id: "Console view"
#   name: "CI Builders"
#   repo_url: "https://noop.com"
#   refs: "regexp:refs/tags/blah"
#   refs: "regexp:refs/branch-heads/\\d+\\.\\d+"
#   exclude_ref: "refs/heads/master"
#   manifest_name: "REVISION"
#   builders: <
#     name: "buildbot/master/builder"
#     name: "buildbucket/luci.infra.ci/linux ci builder"
#     category: "a|b"
#     short_name: "lnx"
#   >
#   builders: <
#     name: "buildbucket/luci.infra.ci/cron builder"
#     category: "cron"
#   >
#   builders: <
#     name: "buildbucket/luci.infra.inline/triggered builder"
#   >
#   builders: <
#     name: "buildbot/master/very buildbot"
#   >
#   favicon_url: "https://storage.googleapis.com/chrome-infra-public/logo/favicon.ico"
#   header: <
#     links: <
#       name: "a"
#       links: <
#         text: "a"
#       >
#     >
#     links: <
#       name: "b"
#       links: <
#         text: "b"
#       >
#     >
#   >
#   include_experimental_builds: true
# >
# logo_url: "https://storage.googleapis.com/chrome-infra-public/logo/chrome-infra-logo-200x200.png"
# build_bug_template: <
#   summary: "Bug summary"
#   description: "Everything is broken"
#   monorail_project: "tutu, all aboard"
#   components: "Stuff>Hard"
# >
# ===
#
# === luci-scheduler.cfg
# job: <
#   id: "cron builder"
#   schedule: "0 6 * * *"
#   acl_sets: "ci"
#   buildbucket: <
#     server: "cr-buildbucket.appspot.com"
#     bucket: "ci"
#     builder: "cron builder"
#   >
# >
# job: <
#   id: "generically named builder"
#   acls: <
#     role: TRIGGERER
#     granted_to: "builder@example.com"
#   >
#   acl_sets: "ci"
#   buildbucket: <
#     server: "cr-buildbucket.appspot.com"
#     bucket: "ci"
#     builder: "generically named builder"
#   >
# >
# job: <
#   id: "linux ci builder"
#   acl_sets: "ci"
#   triggering_policy: <
#     kind: GREEDY_BATCHING
#     max_concurrent_invocations: 5
#     max_batch_size: 10
#   >
#   buildbucket: <
#     server: "cr-buildbucket.appspot.com"
#     bucket: "ci"
#     builder: "linux ci builder"
#   >
# >
# job: <
#   id: "triggered builder"
#   acls: <
#     role: TRIGGERER
#     granted_to: "builder@example.com"
#   >
#   acl_sets: "inline"
#   buildbucket: <
#     server: "cr-buildbucket.appspot.com"
#     bucket: "inline"
#     builder: "triggered builder"
#   >
# >
# job: <
#   id: "triggerer builder"
#   acl_sets: "inline"
#   buildbucket: <
#     server: "cr-buildbucket.appspot.com"
#     bucket: "inline"
#     builder: "triggerer builder"
#   >
# >
# trigger: <
#   id: "inline poller"
#   schedule: "with 10s interval"
#   acl_sets: "inline"
#   triggers: "triggerer builder"
#   gitiles: <
#     repo: "https://noop.com"
#     refs: "regexp:refs/heads/master"
#     refs: "regexp:refs/tags/blah"
#     refs: "regexp:refs/branch-heads/\\d+\\.\\d+"
#   >
# >
# trigger: <
#   id: "master-poller"
#   schedule: "with 10s interval"
#   acl_sets: "ci"
#   triggers: "generically named builder"
#   triggers: "linux ci builder"
#   gitiles: <
#     repo: "https://noop.com"
#     refs: "regexp:refs/heads/master"
#     refs: "regexp:refs/tags/blah"
#     refs: "regexp:refs/branch-heads/\\d+\\.\\d+"
#   >
# >
# acl_sets: <
#   name: "ci"
#   acls: <
#     role: OWNER
#     granted_to: "group:admins"
#   >
#   acls: <
#     granted_to: "group:all"
#   >
#   acls: <
#     role: TRIGGERER
#     granted_to: "group:devs"
#   >
# >
# acl_sets: <
#   name: "inline"
#   acls: <
#     role: OWNER
#     granted_to: "group:admins"
#   >
#   acls: <
#     granted_to: "group:all"
#   >
# >
# ===
#
# === project.cfg
# name: "infra"
# access: "group:all"
# ===
