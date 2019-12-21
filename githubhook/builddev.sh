set -x

echo "URL"
echo $1

echo "Branch"
echo $2


cd /
git clone $1 TalesFromTheRimDev

cd /TalesFromTheRimDev
git checkout $2

git pull


./configure
cd src
make

rm -rf /TalesFromTheRimDev/bin/empiredev

cp /TalesFromTheRimDev/bin/empire /TalesFromTheRimDev/bin/empiredev


pkill empiredev
