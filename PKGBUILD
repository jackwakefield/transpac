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
source=("$url/archive/$pkgname-$pkgver.tar.gz")
sha256sums=('9785c22f970e1047b76ce16155be2ed560606cfed027c9ac8fef8aa158ae9312')

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