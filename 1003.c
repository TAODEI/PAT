#include<stdio.h>
#include<string.h>
void main(){
    int n;
    scanf("%d",&n);
    for(int i=0;i<n;i++){
        char a[101]={0};
        scanf("%s",a);
        int x=0,j,y,h=0;
        for(j=0;j<strlen(a);j++){
            if(a[j]!='A'&&a[j]!='P'&&a[j]!='T'){
                puts("NO");
				h=1;
                break;
            }
        }
		if(h==1) continue;
        for(j=0;j<strlen(a);j++){
            if(a[j]=='T'){
                puts("NO");
                h=1;
				break;
            }
            if(a[j]=='P') break;
            if(a[j]=='A'&&a[j+1]=='P'){
                x=j+1;
                break;
            }
        }
		if(h==1) continue;
        if(a[x+1]!='A'||(a[x+1]=='A'&&a[x+2]=='P')||(a[x+1]=='A'&&a[x+2]=='A'&&a[x+3]!='T')){
            puts("NO");
            continue;
        }
        if(a[x+2]=='T'){
           for(j=x+3;j<strlen(a);j++){
               if(a[j]!='A'){
                    puts("NO");
                    h=1;
					break;
               }
           }
            y=strlen(a)-x-3;
        }
		if(h==1) continue;
        if(a[x+3]=='T'){
           for(j=x+4;j<strlen(a);j++){
               if(a[j]!='A'){
                    puts("NO");
                    h=1;
					break;
               }
           }
            y=strlen(a)-x-4;
        }
		if(h==1) continue;
        //printf("YES%d %d   ",i,y);
        if(y==x||y==2*x) printf("YES\n");
    }
}
