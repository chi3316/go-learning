package variable;

public class Test {
    public static void main(String[] args) {
        int a = 101;
        if (a > 100) {
            // int a = 99; //Duplicate local variable a
            System.out.println(a);
            int[] res = new int[2];
            int[] res2 = new int[] { 0, 1 };
            System.out.println(res);
            System.out.println(res2);
        }
    }
}
