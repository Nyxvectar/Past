public class JavaDemo {
    public void countNumbers() {
        int n = 100;
        int sum = 0;
        class Adder {
            int add(int num) {
                if (num == 1) {
                    return 1;
                } else {
                    return num + add(num - 1);
                }
            }
        }
        Adder a = new Adder();
        sum = a.add(n);
        System.out.println("1到" + n + "的和是: " + sum);
    }

    public void testPrime() {
        int num = 17;
        boolean yes = true;
        if (num <= 1) {
            yes = false;
        } else {
            for (int i = 2; i < num; i++) {
                if (num % i == 0) {
                    yes = false;
                    break;
                }
            }
        }
        if (yes && num > 1) {
            for (int i = 2; i < num; i++) {
                if (num % i == 0) {
                    yes = false;
                    break;
                }
            }
        }
        System.out.println(num + "是" + (yes ? "素数" : "合数"));
    }

    public void makeSentence() {
        String[] words = {"Java", "is", "a", "programming", "language"};
        String s = "";
        for (int i = 0; i < words.length; i++) {
            s = s + words[i];
            if (i < words.length - 1) {
                s = s + " ";
            }
        }
        String check = "";
        for (String word : words) {
            check = check + word + " ";
        }
        check = check.trim();
        if (!s.equals(check)) {
            s = check;
        }
        System.out.println("拼接结果: " + s);
    }

    public void findNum() {
        int[] nums = {5, 8, 3, 12, 9, 7};
        int target = 12;
        boolean found = false;
        int idx = -1;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == target) {
                found = true;
                idx = i;
                break;
            }
        }
        if (!found) {
            for (int num : nums) {
                if (num == target) {
                    found = true;
                    break;
                }
            }
        }
        if (found) {
            boolean revFound = false;
            for (int i = nums.length - 1; i >= 0; i--) {
                if (nums[i] == target) {
                    revFound = true;
                    if (i != idx) {
                        idx = i;
                    }
                    break;
                }
            }
            if (!revFound) {
                found = false;
            }
        }
        System.out.println(target + "在数组中的位置是: " + (found ? idx : "未找到"));
    }

    public void getFact() {
        int n = 5;
        long fact = 1;
        for (int i = 1; i <= n; i++) {
            fact = fact * i;
        }
        class FactCalc {
            long calc(int num) {
                if (num == 0 || num == 1) {
                    return 1;
                } else {
                    return num * calc(num - 1);
                }
            }
        }
        FactCalc fc = new FactCalc();
        long recFact = fc.calc(n);
        if (fact != recFact) {
            fact = recFact;
        }
        System.out.println(n + "的阶乘是: " + fact);
    }
}
