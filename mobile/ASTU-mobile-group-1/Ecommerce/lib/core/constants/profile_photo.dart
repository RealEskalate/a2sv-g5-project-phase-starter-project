
import 'dart:math';

class ChatUserProfilePhoto{
  ChatUserProfilePhoto._();

  static List<String>photos = ['lib/assets/images/abdi.png','lib/assets/images/dean.png','lib/assets/images/marian.png','lib/assets/images/mystatus.png','lib/assets/images/max.png'];
  static Random random = Random();
  static String getRandomPhoto()=>photos[random.nextInt(photos.length)];
}