// 2021-07-21

#include<stdio.h>
int main()
{
    int a,b,c,d,e,f;
    scanf("%d%d%d%d",&a,&b,&c,&d);
    e = a*60+b;
    f = c*60+d;
    printf("%d %d",(f-e)/60,(f-e)%60);
    return 0;
}

// Passed.
