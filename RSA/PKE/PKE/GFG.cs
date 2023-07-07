using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace PKE
{
	
	public class GFG
	{

		public static int __gcd(int a, int b)
		{
			if (a == 0 || b == 0)
				return 0;

			if (a == b)
				return a;

			if (a > b)
				return __gcd(a - b, b);

			return __gcd(a, b - a);
		}

		public static void coprime(int a, int b)
		{

			if (__gcd(a, b) == 1)
				MessageBox.Show("Co-Prime");
				
			else
				MessageBox.Show("Not Co-Prime");
		
		}		
	}
}
