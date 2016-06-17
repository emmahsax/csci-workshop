public class VerySecureEncryption {
	public String encrypt(String message, int[] key, int K) {
		if (K == 0) {
			return message;
		}
		
		String scrambledMessage = new String();
		
		for (int j = 0; j < K; j++) { //mix up k times
			for (int w = 0; w < key.length; w++)  { //go through 0 through n
				for (int i = 0; i < key.length; i++) { //go through entire array key
					if (key[i] == w) {
						scrambledMessage += message.charAt(i);
					}
				
				}
			}
			message = scrambledMessage;
			scrambledMessage = "";
		}
		
		return message;
	}
}

