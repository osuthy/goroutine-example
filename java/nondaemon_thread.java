class Main {

  public static void main(String[] args) {
    new Thread() {
      public void run() {
        while(true) {
          System.out.println("死ななーい");
        }
      }
    }.start();
  }
}
