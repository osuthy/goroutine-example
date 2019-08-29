class Main {

  public static void main(String[] args) {
    Object lock1 = new Object();
    Object lock2 = new Object();

    new Thread() {
      public void run() {
        synchronized(lock1) {
          try {
            Thread.sleep(2);
          } catch(Exception e) {}
          synchronized(lock2) {
            System.out.println("AHOOOOOOOO");
          }
        }
      }
    }.start();
    new Thread() {
      public void run() {
        synchronized(lock2) {
          synchronized(lock1) {
            System.out.println("BAKAAAAAAA");
          }
        }
      }
    }.start();
  }
}
