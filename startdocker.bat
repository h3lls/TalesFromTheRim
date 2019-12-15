REM Edit the path location to point to the folder
docker build -t tales . 
docker run -p 5000:5000 -v %cd%:/TalesFromTheRim -it tales "/dockerbuild.sh"
