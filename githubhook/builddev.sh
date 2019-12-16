echo "URL"
echo $1

echo "Branch"
echo $2

#rm -rf /build/*
#cd /build
cd /
git clone $1 TalesFromTheRimDev

cd /TalesFromTheRimDev
git checkout $2

git pull

#cd /build/TalesFromTheRim/
./configure
cd src
make

rm -rf /TalesFromTheRimDev/bin/empiredev

cp /TalesFromTheRimDev/bin/empire /TalesFromTheRimDev/bin/empiredev
#mkdir -p /TalesFromTheRim/bin/ && cp ../bin/empire /TalesFromTheRimDev/bin/empiredev

pkill empiredev
