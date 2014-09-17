# Maintainer: Jack Wakefield <jackwakefield91@gmail.com>

pkgname=transpac
pkgver=1.0.0
pkgrel=1
pkgdesc="A transparent proxy which uses proxy auto-config (PAC) files for forwarding"
arch=('x86_64' 'i686')
url="https://github.com/jackwakefield/$pkgname/"
license=('Apache')
makedepends=('go')
options=('!strip' '!emptydirs')
source=("$url/releases/download/v$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('18a6d222110863045f8790c58831e8e35d627a0b9c5562dde475fba6bdb8fc8e')

build() {
  cd "$pkgname-$pkgver"

  make install_deps
  make
}

package() {
  cd "$pkgname-$pkgver"

  install -Dm755 "$pkgname-$pkgver" "$pkgdir/usr/bin/$pkgname"
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}