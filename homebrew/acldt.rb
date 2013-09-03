require "formula"

class Acldt < Formula
  VERSION = "0.3.0"
  ARCH = if MacOS.prefer_64_bit?
           "amd64"
         else
           "386"
         end

  homepage "https://github.com/acl-services/acldt"
  head "https://github.com/acl-services/acldt.git"
  url "https://github.com/acl-services/acldt/releases/download/v#{VERSION}/acldt_#{VERSION}-snapshot_darwin_#{ARCH}.tar.gz"
  version VERSION

  def install
    bin.install "acldt"
  end

  def caveats; <<-EOS.undent
  To upgrade gh, run `brew upgrade https://raw.github.com/acl-services/acldt/master/homebrew/acldt.rb`

  More information here: https://github.com/acl-services/acldt/blob/master/README.md
    EOS
  end

  test do
    assert_equal VERSION, `#{bin}/acldt version`.split.last
  end
end
