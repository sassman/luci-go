def gen1(ctx):
  fail("gen1 failed")
lucicfg.generator(impl = gen1)

trace = stacktrace()

def gen2(ctx):
  fail("gen2 failed", trace=trace)
lucicfg.generator(impl = gen2)

def gen3(ctx):
  fail("dedupped error", trace=trace)
lucicfg.generator(impl = gen3)

def gen4(ctx):
  fail("dedupped error", trace=trace)
lucicfg.generator(impl = gen4)

# Expect errors:
#
# Traceback (most recent call last):
#   //testdata/misc/generator_errors.star:2: in gen1
# Error: gen1 failed
#
# Traceback (most recent call last):
#   //testdata/misc/generator_errors.star:5: in <toplevel>
# Error: gen2 failed
#
# Traceback (most recent call last):
#   //testdata/misc/generator_errors.star:5: in <toplevel>
# Error: dedupped error
