import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Main {
    
    static int n;
    static int k;
    static int[] arr;
    static int ans;

    public static void main(String[] args) throws IOException {
         BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
         StringTokenizer st = new StringTokenizer(br.readLine());
         
         n = Integer.parseInt(st.nextToken());
         k = Integer.parseInt(st.nextToken());
        
        if (k < 5) {
            System.out.println(0);
            return;
        }
         
         arr = new int[n];
         
         for (int i = 0; i < n; i++) {
        	 int res = 0;
        	 char[] crr = br.readLine().toCharArray();
        	 for (int j = 0; j < crr.length; j++) {
        		 
        		 if ((res & (1 << (crr[j] - 'a'))) != 0) continue;
        		 
        		 res |= 1 << (crr[j] - 'a');
        	 }
        	 
        	 arr[i] = res;
         }

         ans = 0;
         
         int prob = 0;
         prob |= 1 << ('a' - 'a');
         prob |= 1 << ('n' - 'a');
         prob |= 1 << ('t' - 'a');
         prob |= 1 << ('i' - 'a');
         prob |= 1 << ('c' - 'a');
         
         comb(0, 5, prob);
         System.out.println(ans);
    }
    static void comb(int depth, int cnt, int prob) {
    	
    	if (26 - depth < k - cnt) return;
    	
    	if (cnt == k) {
    		
    		int res = 0;
    		
    		for (int i = 0; i < n; i++) {
    			if ((prob & arr[i]) == arr[i]) res++;
    			
    			if (res + n - i - 1 <= ans) return;
    		}
    		
    		if (ans < res) ans = res;
    		return;
    	}
    	
    	if (depth == 26) return;
    	
    	if (depth == 'a' - 'a' || depth == 'n' - 'a' || depth == 't' - 'a' || depth == 'i' - 'a' || depth == 'c' - 'a') {
    		comb(depth + 1, cnt, prob);
    	} else {
    		comb(depth + 1, cnt + 1, prob | 1 << depth);
        	comb(depth + 1, cnt, prob);
    	}
    }
}