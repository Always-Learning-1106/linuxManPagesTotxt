This program will put any linux man commands into folders named after the command, as well as a text file inside the folder with the same name

In order for this to work you must have a text file named commands.txt inside the same folder as this program. 

The program reads line by line off of the commands.txt file and each individual command one by one, so only put one command per line. EXAMPLE below
inside of commands.txt...
ls
cd
mkdir
mv
cp
wget
-----------
put as many commands as you want, one by one, one per line, and the program will copy the contents of the man page for that command
and put it into a .txt file sharing the same name as the command, which will be in a folder with the same name as the command. for instance see below example
command "ls" from first line of commands.txt
folder = ls
file inside folder = ls.txt
content inside ls.txt = man ls results as you would see in bash terminal

SIDE NOTE
don't type man before the command, just type the commands line by line inside of the commands.txt file as the example shows 

BONUS
Make sure you spell the commands correctly because I am lazy and didn't do error checks. 
I think it just makes files with .txt files inside that have the mispelled command names, so you could use this as a file generator I suppose.
Write a bunch of lines that aren't linux commands of filenames you want to make and bang you can generate files on a txt document. 

Works on linux. Never tested any other OS because it was just a side project for a few hours. 