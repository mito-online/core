# MITO Core

[Swagger Specs](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/mito-online/core/master/domain/public.swagger.json)

[core Package](https://buf.build/mito/core)

To compile:

```sh
buf generate proto
```

To publish in [buf.build](https://buf.build/):

```sh
cd proto
buf push
```

More actions: checkout `Makefile`