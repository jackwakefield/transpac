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
source=("$url/archive/v$pkgver.tar.gz")
sha256sums=('582895d391a7e7cac901bd256dd9bc8f644141e873107b283a7c915e961c0493')

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