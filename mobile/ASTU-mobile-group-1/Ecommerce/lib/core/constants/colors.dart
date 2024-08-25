import 'dart:math';

import '../../features/auth/presentation/pages/pages.dart';

class ChatColors{
  ChatColors._();
  static List<Color>profileColors = [const Color.fromARGB(255,197, 197, 197),const Color.fromARGB(255,245, 183, 190),const Color.fromARGB(255,199, 254, 224),const Color.fromARGB(255,200, 236, 253),const Color.fromARGB(255,209, 150, 12)];

  static Random random = Random();
  static Color getRandomColor()=>profileColors[random.nextInt(profileColors.length)];

  static Color lightBlueColor  = Colors.blue.shade500;
  static Color mediumBlueColor  = Colors.blue.shade800;
   static Color recievedColor  = const Color.fromARGB(255, 174, 189, 202);
}