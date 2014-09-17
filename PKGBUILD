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
install=transpac.install
source=("$url/releases/download/v$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('eadddfc963d3e937f6cccc51bd5314c9c2143c8ebdb22d92211969f70703e604')

build() {
  cd "$pkgname-$pkgver"

  make install_deps
  make
}

package() {
  cd "$pkgname-$pkgver"

  install -Dm755 "$pkgname-$pkgver" "$pkgdir/usr/bin/$pkgname"
  install -Dm644 config.toml "$pkgdir/etc/$pkgname/config.toml"
  install -Dm644 transpac.service "$pkgdir/usr/lib/systemd/system/transpac.service"
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}