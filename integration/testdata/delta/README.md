# Test data for deltas

## Data files

The data files we use to populate our test Docker images are divided in two
groups. The first group contains random data:

* `000.data`: An empty file.
* `001.data`: Another empty file.
* `002.data`: 256 bytes of random data.
* `003.data`: 333 bytes of random data.
* `004.data`: 1024 bytes of random data.
* `005.data`: 1024 bytes of random data (different from 003).
* `006.data`: 3307 bytes of random data.
* `007.data`: 8 KiB of random data.
* `008.data`: 9876 bytes of random data.
* `009.data`: 10 KiB of random data.
* `010.data`: 128 KiB of random data.
* `011.data`: 150000 bytes of random data.
* `012.data`: 500003 bytes of random data.
* `013.data`: 1000000 bytes of random data.
* `014.data`: 1000000 bytes of random data (different from 012).
* `015.data`: 1 MiB of random data.
* `016.data`: 1 MiB of random data (different from 015).

The second group is made by concatenating the files from the first group. In the
descriptions below, we use the _001 + 002 + 010_ notation to mean "the
concatenation of `001.data`, `002.data`, and `010.data`." We say _3 * 007_ as a
shortcut to "the contents of `007.data` repeated three times." And we can go
fancy and use _4 * (002 + 004)_ to mean "the concatenation of `002.data` and
`004.data` repeated four times."

* `100.data`: 015 + 013 + 012 + 014 + 010.
* `101.data`: 016 + 013 + 012 + 014 + 010.
