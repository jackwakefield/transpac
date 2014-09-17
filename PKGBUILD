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
sha256sums=('fe1d402648631ecee06e74c7831003727acc135db5c78186b5e276b038972bf5')

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