using System;

namespace class_csharp
{
    class Program
    {
        static void Main(string[] args)
        {
            var foo = 1;
            mutate(ref foo);
            Console.WriteLine(foo.ToString());
        }

        static void mutate(ref int val) 
        {
            val = 0;
        }
    }
}
