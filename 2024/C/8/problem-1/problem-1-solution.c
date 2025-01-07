#include <stdio.h>
#include <stdlib.h>
#include <string.h>


int main (int argc, char * argv[])
{
    char *filename = "inputdata.txt";
    long result=0;
    int part=1;
    char buf[55];
    char grid[55][55]={0};
    char nodes[55][55]={0};
    int size=0;
    FILE * fp;
    if(argc > 1)
        part=atoi(argv[1]);
    if(argc > 2)
        filename=argv[2];
    if((fp = fopen (filename, "r"))==NULL)
    {
        printf("error opening file %s\n", filename);
        return 1;
    }
    while(fgets(buf, 55, fp) != NULL)
    {
        strncpy(grid[size++],buf,sizeof(grid[size]));
    }
    for(int x=0;x<size;x++)
        for(int y=0;y<size;y++)
        {
            if(grid[x][y]!='.')
            for(int x2=x;x2<size;x2++)
                for(int y2=0;y2<size;y2++)
                {
                    if(x2==x && y2<= y)
                        continue;
                    if(grid[x][y]== grid[x2][y2])
                    {
                        int slopex=x2-x;
                        int slopey=y2-y;
                        if(part==1)
                        {
                            if((x-slopex<size)&&(x-slopex>=0)&&(y-slopey<size)&&(y-slopey>=0))
                                nodes[x-slopex][y-slopey]=1;
                            if((x2+slopex<size)&&(x2+slopex>=0)&&(y2+slopey<size)&&(y2+slopey>=0))
                                nodes[x2+slopex][y2+slopey]=1;
                        }
                        else
                        {
                            for(int i=0;(x-(slopex*i)<size)&&(x-(i*slopex)>=0)&&(y-(i*slopey)<size)&&(y-(i*slopey)>=0);i++)
                                nodes[x-(i*slopex)][y-(i*slopey)]=1;
                            for(int i=1;(x+(slopex*i)<size)&&(x+(i*slopex)>=0)&&(y+(i*slopey)<size)&&(y+(i*slopey)>=0);i++)
                                nodes[x+(i*slopex)][y+(i*slopey)]=1;
                        }
                    }
                }
        }
    for(int x=0;x<size;x++)
        for(int y=0;y<size;y++)
            result+=nodes[x][y];
    printf("%ld\n",result);
    return 0;
}