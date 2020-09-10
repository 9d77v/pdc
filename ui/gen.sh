
pbjs -t static-module -w commonjs -o compiled.js ./proto/*.proto
pbts -o compiled.d.ts compiled.js