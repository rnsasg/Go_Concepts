#include <bits/stdc++.h>

using namespace std;
int main()
{
    // Initializing a vector
    vector<pair<int,int>> v1 = { {1,2}, {4,3},{5,6},{7,8},{0,1} };
 
    // Converting vector into a heap
    // using make_heap()
    make_heap(v1.begin(), v1.end());
 
    // Displaying the maximum element of heap
    // using front()
    cout << "The maximum element of heap is : ";
   cout << v1[0].first<<v1[0].second << endl;
 
    return 0;
}

// comparator function to make min heap 

class Solution {
public: 
    vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        
    // Initializing a vector
    vector<pair<int,int>> heap;
    vector<vector<int>> res;
    pair<int,int> ele;
    vector<int> temp;
    struct greaters{ 
    bool operator()(const long& a,const long& b) const{ 
        return a>b; 
    } 
    };
        
    for (int i=0;i< nums1.size();i++){
        for (int j = 0;j<nums2.size();j++){
            heap.push_back(make_pair(nums1[i],nums2[j]));
        }
    }
 
    make_heap(heap.begin(), heap.end(),greaters());
    int t = 0;
    for(int i = 0 ;t<k &&  t < heap.size();i++){
        ele = heap.front();
        temp.push_back(ele.first);
        temp.push_back(ele.second);
        res.push_back(temp);
        temp.clear();
        t++;
    }
    return res;
    }
};