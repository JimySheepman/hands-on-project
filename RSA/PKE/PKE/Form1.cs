using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace PKE
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void label1_Click(object sender, EventArgs e)
        {

        }

        private void label2_Click(object sender, EventArgs e)
        {

        }

        private void checkBox1_CheckedChanged(object sender, EventArgs e)
        {
            if (checkBox1.Checked)
            {
                groupBox1.Enabled = true;
                groupBox1.Visible = true;

            }
            else
            {
                groupBox1.Enabled = false;
                groupBox1.Visible = false;
            }
        }

        private void label15_Click(object sender, EventArgs e)
        {

        }

        private void button4_Click(object sender, EventArgs e)
        {
            //c=(m^e)mod n
            int numpe = Convert.ToInt32(textBox10.Text);  //e
            int numpn = Convert.ToInt32(textBox11.Text);  //n       
            String s = (textBox12.Text);
            for (int i = 0; i < s.Length; i++)
            {
                int m = (int)s[i];
                textBox13.AppendText(m.ToString() + ',');
                int c = 1;
                for (int k = 0; k < numpe; k++)
                {
                    c = (c * m) % numpn;

                }
                c = c % numpn;
                textBox14.AppendText(c.ToString("X4"));
            }



            /*
            ASCIIEncoding encoding = new ASCIIEncoding();
            Byte[] bytes = encoding.GetBytes(s);
            for (int i = 0; i < bytes.Length; i++)
            {
                int m = bytes[i];
                int c = 1;
                for (int k = 0; k < numpe; k++)
                {
                    c = (c * m) % numpn;

                }
                c = c % numpn;
                bytes[i] = Convert.ToByte(c);
            }
            foreach (Byte value in bytes)
            {
                textBox13.AppendText(value.ToString() + ',');
            }
            foreach (Byte value in bytes)
            {
                textBox14.AppendText(value.ToString("X4"));

            }
            */
        }

        private void groupBox1_Enter(object sender, EventArgs e)
        {

        }

        private void checkBox2_CheckedChanged(object sender, EventArgs e)
        {
            if (checkBox2.Checked)
            {
                groupBox2.Enabled = true;
                groupBox2.Visible = true;

            }
            else
            {
                groupBox2.Enabled = false;
                groupBox2.Visible = false;
            }
        }

        private void checkBox3_CheckedChanged(object sender, EventArgs e)
        {
            if (checkBox3.Checked)
            {
                groupBox3.Enabled = true;
                groupBox3.Visible = true;

            }
            else
            {
                groupBox3.Enabled = false;
                groupBox3.Visible = false;
            }

        }

        private void textBox1_TextChanged(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            // hata mesajından sonra hesaplama yapıyor düzeltilmesi gerek!!
            // e<n kontrol edilmesi
            int count = 0;
            int number1 = Convert.ToInt32(textBox1.Text);
            int number2 = Convert.ToInt32(textBox2.Text);

                for (int i = 2; i <= (number1 / 2); i++)
                {
                    if (number1 % i == 0)
                    {
                        count++;
                        MessageBox.Show("This Number Not Prime (p)");
                        
                    }
                }
                for (int i = 2; i <= (number2 / 2); i++)
                {
                    if (number2 % i == 0)
                    {
                        count++;
                        MessageBox.Show("This Number Not Prime (q)");
                        
                    }
                }
                textBox3.Text = (number1 * number2).ToString();
                textBox4.Text = ((number1 - 1) * (number2 - 1)).ToString();
        }

        private void button2_Click(object sender, EventArgs e)
        {
            // hata mesajları eksik
            int nnum = Convert.ToInt32(textBox3.Text);   //   n sayısı
            int eenum = Convert.ToInt32(textBox9.Text);   // e sayısı
            int znum = Convert.ToInt32(textBox4.Text);   //  z sayısı
            int dnum=0;
            int count = 0;
            for (int i = 0; i < znum; i++)
            {
                if (count == 1)
                {
                    dnum = dnum - 1;
                    break;
                }
                else
                {
                    int v = (eenum * dnum);
                    count = v % znum;
                    dnum = i;
                }
            }
            GFG.coprime(znum, eenum);
            textBox5.Text = textBox9.Text;
            textBox7.Text = textBox3.Text;
            textBox8.Text = textBox3.Text;
            textBox6.Text = Convert.ToString(dnum);
        }

        private void textBox13_TextChanged(object sender, EventArgs e)
        {

        }

        private void Form1_Load(object sender, EventArgs e)
        {
          
        }

        private void button9_Click(object sender, EventArgs e)
        {
            int x = 0;
            int x1 = 0;
            int x2 = 0;
            int x3 = 0;
            int mq = 0;
            int numpd = Convert.ToInt32(textBox19.Text); // d
            int numpnn = Convert.ToInt32(textBox18.Text);  // n
            string l = textBox17.Text;
            string a;
            ASCIIEncoding aSCII = new ASCIIEncoding();
            
            Byte[] vs = aSCII.GetBytes(l);
            char[] vs1 = l.ToCharArray();
            

            for (int i = 3; i < vs1.Length; i+=4)
            {
                if (vs1[i - 3]=='0')
                {
                    x = 0;
                }
                if (vs1[i - 3] == '1')
                {
                    x = 1*(16*16*16);
                }
                if (vs1[i - 3] == '2')
                {
                    x = 2 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == '3')
                {
                    x = 3 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == '4')
                {
                    x = 4 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == '5')
                {
                    x = 5 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == '6')
                {
                    x = 6   * (16 * 16 * 16);
                }
                if (vs1[i - 3] == '7')
                {
                    x = 7 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == '8')
                {
                    x = 8 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == '9')
                {
                    x = 9 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == 'A')
                {
                    x = 10 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == 'B')
                {
                    x = 11 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == 'C')
                {
                    x = 12 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == 'D')
                {
                    x = 13 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == 'E')
                {
                    x = 14 * (16 * 16 * 16);
                }
                if (vs1[i - 3] == 'F')
                {
                    x = 15 * (16 * 16 * 16);
                }
                if (vs1[i-2]=='0')
                {
                    x1 = 0;
                }
                if (vs1[i - 2] == '1')
                {
                    x1 = 1*(16*16);
                }
                if (vs1[i - 2] == '2')
                {
                    x1 = 2 * (16 * 16);
                }
                if (vs1[i - 2] == '3')
                {
                    x1 = 3 * (16 * 16);
                }
                if (vs1[i - 2] == '4')
                {
                    x1 = 4 * (16 * 16);
                }
                if (vs1[i - 2] == '5')
                {
                    x1 = 5 * (16 * 16);
                }
                if (vs1[i - 2] == '6')
                {
                    x1 = 6 * (16 * 16);
                }
                if (vs1[i - 2] == '7')
                {
                    x1 = 7 * (16 * 16);
                }
                if (vs1[i - 2] == '8')
                {
                    x1 = 8 * (16 * 16);
                }
                if (vs1[i - 2] == '9')
                {
                    x1 = 9 * (16 * 16);
                }
                if (vs1[i - 2] == 'A')
                {
                    x1 = 10 * (16 * 16);
                }
                if (vs1[i - 2] == 'B')
                {
                    x1 = 11 * (16 * 16);
                }
                if (vs1[i - 2] == 'C')
                {
                    x1 = 12* (16 * 16);
                }
                if (vs1[i - 2] == 'D')
                {
                    x1 = 13 * (16 * 16);
                }
                if (vs1[i - 2] == 'E')
                {
                    x1 = 14 * (16 * 16);
                }
                if (vs1[i - 2] == 'F')
                {
                    x1 = 15 * (16 * 16);
                }
                if (vs1[i-1]=='0')
                {
                    x2 = 0;
                }
                if (vs1[i - 1] == '1')
                {
                    x2 = 1*16;
                }
                if (vs1[i - 1] == '2')
                {
                    x2 = 2 * 16;
                }
                if (vs1[i - 1] == '3')
                {
                    x2 = 3 * 16;
                }
                if (vs1[i - 1] == '4')
                {
                    x2 = 4 * 16;
                }
                if (vs1[i - 1] == '5')
                {
                    x2 = 5 * 16;
                }
                if (vs1[i - 1] == '6')
                {
                    x2 = 6 * 16;
                }
                if (vs1[i - 1] == '7')
                {
                    x2 = 7 * 16;
                }
                if (vs1[i - 1] == '8')
                {
                    x2 = 8* 16;
                }
                if (vs1[i - 1] == '9')
                {
                    x2 = 9 * 16;
                }
                if (vs1[i - 1] == 'A')
                {
                    x2 = 10 * 16;
                }
                if (vs1[i - 1] == 'B')
                {
                    x2 = 11 * 16;
                }
                if (vs1[i - 1] == 'C')
                {
                    x2 = 12 * 16;
                }
                if (vs1[i - 1] == 'D')
                {
                    x2 = 13 * 16;
                }
                if (vs1[i - 1] == 'E')
                {
                    x2 = 14 * 16;
                }
                if (vs1[i - 1] == 'F')
                {
                    x2 = 15 * 16;
                }
                if (vs1[i]=='0')
                {
                    x3 = 0;
                }
                if (vs1[i] == '1')
                {
                    x3 = 1;
                }
                if (vs1[i] == '2')
                {
                    x3 = 2;
                }
                if (vs1[i] == '3')
                {
                    x3 = 3;
                }
                if (vs1[i] == '4')
                {
                    x3 = 4;
                }
                if (vs1[i] == '5')
                {
                    x3 = 5;
                }
                if (vs1[i] == '6')
                {
                    x3 = 6;
                }
                if (vs1[i] == '7')
                {
                    x3 = 7;
                }
                if (vs1[i] == '8')
                {
                    x3 = 8;
                }
                if (vs1[i] == '9')
                {
                    x3 = 9;
                }
                if (vs1[i] == 'A')
                {
                    x3 = 10;
                }
                if (vs1[i] == 'B')
                {
                    x3 = 11;
                }
                if (vs1[i] == 'C')
                {
                    x3 = 12;
                }
                if (vs1[i] == 'D')
                {
                    x3 = 13;
                }
                if (vs1[i] == 'E')
                {
                    x3 = 14;
                }
                if (vs1[i] == 'F')
                {
                    x3 = 15;
                }

                mq = x3 + x2 + x1 + x;
                textBox16.AppendText(mq.ToString() + ',');
                int c = 1;
                for (int k = 0; k < numpd; k++)
                {
                    c = (c * mq) % numpnn;

                }
                c = c % numpnn;
                //textBox15.AppendText(c.ToString() + ',');
                char ch = (char)c;
               
                textBox15.AppendText(Convert.ToString(ch));
                mq = 0;
            }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            textBox1.Clear();
            textBox2.Clear();
            textBox3.Clear();
            textBox4.Clear();
            textBox5.Clear();
            textBox6.Clear();
            textBox7.Clear();
            textBox8.Clear();
            textBox9.Clear();
        }

        private void button7_Click(object sender, EventArgs e)
        {
            textBox10.Clear();
            textBox11.Clear();
            textBox12.Clear();
            textBox13.Clear();
            textBox14.Clear();
        }

        private void button8_Click(object sender, EventArgs e)
        {
            textBox15.Clear();
            textBox16.Clear();
            textBox17.Clear();
            textBox18.Clear();
            textBox19.Clear();
        }

        private void textBox16_TextChanged(object sender, EventArgs e)
        {

        }

        private void textBox2_TextChanged(object sender, EventArgs e)
        {

        }

        private void button5_Click(Object sender, EventArgs e)
        {
            textBox10.Text = textBox5.Text;
            textBox11.Text = textBox7.Text;
            textBox19.Text = textBox6.Text;
            textBox18.Text = textBox8.Text;
           
        }

        private void button6_Click(Object sender, EventArgs e)
        {
            textBox17.Text = textBox14.Text;
        }
    }
}

