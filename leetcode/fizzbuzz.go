package main

import "fmt"

func main() {
	flag := false
	for i := 1; i <= 100; i++ {
		flag = false
		if i%3 == 0 {
			fmt.Printf("fizz")
			flag = true
		}
		if i%5 == 0 {
			fmt.Printf("buz")
			flag = true
		}
		if flag == false {
			//fmt.Println()
			fmt.Print(i)
		}
		fmt.Printf("\n")
	}
}

Pick pizza slice of size 4, Pepper and Tony will pick slice with size 3 and 5 respectively. Then Pick slices with size 6, finally Pepper and Tony will pick slice of size 2 and 1 respectively. Total = 4 + 6 = 10

class Solution {
	public:
		int A[501];
		int dp[501][501],n,sizeee;
		int make(int i,int j,int cnt)
		{
			if(i>j)
				return 0;
			if(cnt==sizeee)
				return 0;
			if(dp[i][cnt]!=-1)
				return dp[i][cnt];
		  int p=make(i+2,j,cnt+1)+A[i];
			int q=make(i+1,j,cnt);
			
			return dp[i][cnt]=max(p,q);
		}
		int maxSizeSlices(vector<int>& slices) {
			n=slices.size();
			sizeee=n/3;
			memset(dp,-1,sizeof(dp));
			for(int i=0;i<slices.size();i++)
			{
				A[i]=slices[i];
			}
		   int p=make(0,n-2,0);
			 memset(dp,-1,sizeof(dp));
			int q=make(1,n-1,0);
			return max(p,q);
		}
	};
	