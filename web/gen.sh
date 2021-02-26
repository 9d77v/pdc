pbjs -t static-module -w commonjs -o compiled.js ./proto/*.proto
pbts -o compiled.d.ts compiled.js
mv compiled.js ./src/pb/compiled.js
mv compiled.d.ts ./src/pb/compiled.d.ts
