class Main {

  public static void main(String[] args) {
    Thread th = new Thread() {
      public void run() {
        while(true) {
          System.out.println("死ぬよー");
        }
      }
    };
    th.setDaemon(true);
    th.start();
  }
}
