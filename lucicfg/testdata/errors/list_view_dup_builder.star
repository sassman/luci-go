luci.project(
    name = 'project',
    buildbucket = 'cr-buildbucket.appspot.com',
    milo = 'luci-milo.appspot.com',
    swarming = 'chromium-swarm.appspot.com',
)

luci.bucket(name = 'ci')

luci.recipe(
    name = 'main/recipe',
    cipd_package = 'recipe/bundles/main',
)

luci.builder(
    name = 'b',
    bucket = 'ci',
    recipe = 'main/recipe',
)

luci.list_view(
    name = 'View',
    entries = [
        'b',
        luci.list_view_entry('b'),
    ],
)

# Expect errors like:
#
# Traceback (most recent call last):
#   //testdata/errors/list_view_dup_builder.star:25: in <toplevel>
#   ...
# Error: builder luci.builder("ci/b") was already added to luci.list_view("View"), previous declaration:
# Traceback (most recent call last):
#   //testdata/errors/list_view_dup_builder.star:21: in <toplevel>
#   ...
