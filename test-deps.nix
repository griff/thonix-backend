(import ./deps.nix) ++
[
  {
    goPackagePath = "github.com/davecgh/go-spew";
    fetch = {
      type = "git";
      url = "https://github.com/davecgh/go-spew";
      rev = "6d212800a42e8ab5c146b8ace3490ee17e5225f9";
      sha256 = "01i0n1s4j7khb7n6mz2wymniz37q0vbzkgfv7rbi6p9hpg227q93";
    };
  }
  {
    goPackagePath = "github.com/pmezard/go-difflib";
    fetch = {
      type = "git";
      url = "https://github.com/pmezard/go-difflib";
      rev = "d8ed2627bdf02c080bf22230dbb337003b7aba2d";
      sha256 = "0w1jp4k4zbnrxh3jvh8fgbjgqpf2hg31pbj8fb32kh26px9ldpbs";
    };
  }
  {
    goPackagePath = "github.com/stretchr/testify";
    fetch = {
      type = "git";
      url = "https://github.com/stretchr/testify";
      rev = "69483b4bd14f5845b5a1e55bca19e954e827f1d0";
      sha256 = "11lzrwkdzdd8yyag92akncc008h2f9d1bpc489mxiwp0jrmz4ivb";
    };
  }
]