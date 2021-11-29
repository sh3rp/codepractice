

class FizzBuzz {
    public static String[] searchTerms = new String[]{
        "fizz",
        "buzz",
        "fizzbuzz"
    };

    public static void main(String[] args) {
        String[] testStrings = {
            "fizz buzz fizzbuzz",                      // fizz = 2, buzz = 2, fizzbuzz = 1
            "fizzfizz fizz buzz fizzbuz",              // fizz = 4, buzz = 1, fizzbuzz = 0
            "fizzfizzbuzz fizz buzz fizzbuzz",         // fizz = 4, buzz = 3, fizzbuzz = 2
            "fiz fizz buzzbuzz fizzbuzz buzz fizzfiz", // fizz = 3, buzz = 4, fizzbuzz = 1
        };

        for(String str : testStrings) {
            int[] count = Count(searchTerms,str);
            System.out.printf("%s = %d, %s = %d, %s = %d\n",searchTerms[0], count[0], searchTerms[1], count[1], searchTerms[2], count[2]);
        }
        return;
    }

    public static int[] Count(String[] terms, String str) {
        TermState[] states = new TermState[terms.length];
        for(int i=0;i<terms.length;i++) {
            states[i] = new TermState(terms[i],0,0);
        }

        char[] strChars = str.toCharArray();

        for(int i=0;i<strChars.length;i++) {
            for(int j=0;j<states.length;j++) {
                if(states[j].term[states[j].currentIdx] == strChars[i]) {
                    if(states[j].currentIdx == states[j].term.length-1) {
                        states[j].count++;
                        states[j].currentIdx = 0;
                    } else {
                        states[j].currentIdx++;
                    }
                } else {
                    states[j].currentIdx = 0;
                    if(states[j].term[states[j].currentIdx] == strChars[i]) {
                        states[j].currentIdx++;
                    }

                }
            }
        }

        int[] counts = new int[terms.length];
        for(int i=0;i<states.length;i++) {
            counts[i] = states[i].count;
        }

        return counts;
    }

    static class TermState {
        TermState(String term, int count, int currentIdx) {
            this.term = term.toCharArray();
            this.count = count;
            this.currentIdx = currentIdx;
        }
        char[] term;
        int count;
        int currentIdx;
    }
}