class Lume < Formula
  desc "Fast, safe macOS disk cleanup tool — always moves to Trash, never rm"
  homepage "https://github.com/Tyooughtul/lume"
  version "VERSION_PLACEHOLDER"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/Tyooughtul/lume/releases/download/vVERSION_PLACEHOLDER/lume-darwin-arm64"
      sha256 "SHA256_ARM64_PLACEHOLDER"

      def install
        bin.install "lume-darwin-arm64" => "lume"
      end
    elsif Hardware::CPU.intel?
      url "https://github.com/Tyooughtul/lume/releases/download/vVERSION_PLACEHOLDER/lume-darwin-amd64"
      sha256 "SHA256_AMD64_PLACEHOLDER"

      def install
        bin.install "lume-darwin-amd64" => "lume"
      end
    end
  end

  test do
    assert_match "Lume", shell_output("#{bin}/lume -help")
  end
end
