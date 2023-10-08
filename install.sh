if [ "$(uname)" == "Darwin" ]
then
    curl -L http://tip.ruanfor.com/bin/mac/tip -o /usr/local/bin/tip    
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]
then 
    curl -L http://tip.ruanfor.com/bin/linux/tip -o /usr/local/bin/tip
fi
chmod a+x /usr/local/bin/tip