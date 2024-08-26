import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class MessageBubble extends StatelessWidget {
  final String message;
  final bool isMe;
  final String messageType;

  const MessageBubble({
    required this.message,
    required this.isMe,
    required this.messageType,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Widget messageContent;

    switch (messageType) {
      case 'text':
        messageContent = Text(
          message,
          style: TextStyle(color: isMe ? Colors.white : Colors.black),
        );
        break;
      case 'image':
        messageContent = Image.network(message); // Assuming message is a URL
        break;
      case 'voice':
        messageContent = Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            Icon(Icons.play_arrow),
            const SizedBox(width: 5),
            Text("Voice Message"),
          ],
        );
        break;
      default:
        messageContent = Text(message);
    }

    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: Row(
        mainAxisAlignment:
            isMe ? MainAxisAlignment.end : MainAxisAlignment.start,
        children: [
          if (!isMe) ...[
            Column(
              children: [
                Container(
                  width: 40,
                  height: 40,
                  decoration: const BoxDecoration(
                    shape: BoxShape.circle,
                    color: Colors.blue,
                  ),
                  child: Center(
                    child: Image.asset('images/profile.png'),
                  ),
                ),
                const SizedBox(height: 140),
              ],
            ),
            const SizedBox(width: 10),
          ],
          Flexible(
            child: Column(
              crossAxisAlignment:
                  isMe ? CrossAxisAlignment.end : CrossAxisAlignment.start,
              children: [
                Text(
                  'Today',
                  style: GoogleFonts.dmSans(
                    textStyle: const TextStyle(
                      fontWeight: FontWeight.w500,
                      fontSize: 15,
                      color: Color.fromRGBO(0, 14, 8, 1),
                    ),
                  ),
                ),
                const SizedBox(height: 5),
                Row(
                  mainAxisAlignment:
                      isMe ? MainAxisAlignment.end : MainAxisAlignment.start,
                  children: [
                    if (!isMe) const SizedBox(width: 25),
                    Flexible(
                      child: Material(
                        borderRadius: BorderRadius.only(
                          topLeft: Radius.circular(isMe ? 30 : 0),
                          bottomLeft: const Radius.circular(30),
                          bottomRight: const Radius.circular(30),
                          topRight: Radius.circular(isMe ? 0 : 30),
                        ),
                        elevation: 5.0,
                        color: isMe ? Colors.blue : Colors.grey[300],
                        child: Padding(
                          padding: const EdgeInsets.symmetric(
                              horizontal: 20.0, vertical: 15.0),
                          child: messageContent,
                        ),
                      ),
                    ),
                    if (isMe) const SizedBox(width: 15),
                  ],
                ),
                const SizedBox(height: 5),
                Padding(
                  padding: const EdgeInsets.only(left: 35.0),
                  child: Row(
                    mainAxisAlignment:
                        isMe ? MainAxisAlignment.end : MainAxisAlignment.start,
                    children: [
                      Text(
                        '10:00 AM',
                        style: GoogleFonts.dmSans(
                          textStyle: const TextStyle(
                            fontWeight: FontWeight.w500,
                            fontSize: 12,
                            color: Color.fromRGBO(0, 14, 8, 1),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}


// import 'package:flutter/material.dart';
// import 'package:get/get.dart';
// import 'package:audioplayers/audioplayers.dart';
// import '../controllers/audioController.dart';

// class MessageBubble extends StatelessWidget {
//   final String message;
//   final String messageType;
//   final bool isCurrentUser;
//   final String time;
//   final String? duration; // Added for audio duration
//   final int index;

//   MessageBubble({
//     required this.message,
//     required this.messageType,
//     required this.isCurrentUser,
//     required this.time,
//     this.duration,
//     required this.index,
//   });

//   final AudioController audioController = Get.put(AudioController());

//   @override
//   Widget build(BuildContext context) {
//     return Column(
//       crossAxisAlignment: isCurrentUser ? CrossAxisAlignment.end : CrossAxisAlignment.start,
//       children: [
//         if (messageType == 'text') 
//           _textMessage(),
//         if (messageType == 'audio') 
//           _audioMessage(),
//         // Add other message types (e.g., images) as needed
//         Padding(
//           padding: const EdgeInsets.only(top: 4.0),
//           child: Text(
//             time,
//             style: TextStyle(fontSize: 12, color: Colors.grey),
//           ),
//         ),
//       ],
//     );
//   }

//   Widget _textMessage() {
//     return Container(
//       padding: EdgeInsets.all(10),
//       margin: EdgeInsets.symmetric(vertical: 5),
//       decoration: BoxDecoration(
//         color: isCurrentUser ? Colors.blue : Colors.grey.shade200,
//         borderRadius: BorderRadius.circular(8),
//       ),
//       child: Text(
//         message,
//         style: TextStyle(color: isCurrentUser ? Colors.white : Colors.black),
//       ),
//     );
//   }








//   Widget _audioMessage() {
//     return Container(
//       width: MediaQuery.of(context).size.width * 0.6,
//       padding: EdgeInsets.all(8),
//       margin: EdgeInsets.symmetric(vertical: 5),
//       decoration: BoxDecoration(
//         color: isCurrentUser ? Colors.blue : Colors.grey.shade200,
//         borderRadius: BorderRadius.circular(8),
//       ),
//       child: Row(
//         children: [
//           GestureDetector(
//             onTap: () {
//               audioController.onPressedPlayButton(index, message);
//             },
//             onSecondaryTap: () {
//               audioController.audioPlayer.stop();
//             },
//             child: Obx(
//               () => (audioController.isRecordPlaying && audioController.currentId == index)
//                   ? Icon(Icons.pause, color: isCurrentUser ? Colors.white : Colors.blue)
//                   : Icon(Icons.play_arrow, color: isCurrentUser ? Colors.white : Colors.blue),
//             ),
//           ),
//           SizedBox(width: 8),
//           Expanded(
//             child: Obx(
//               () => LinearProgressIndicator(
//                 backgroundColor: Colors.grey.shade300,
//                 valueColor: AlwaysStoppedAnimation<Color>(isCurrentUser ? Colors.white : Colors.blue),
//                 value: (audioController.isRecordPlaying && audioController.currentId == index)
//                     ? audioController.completedPercentage.value
//                     : 0.0,
//               ),
//             ),
//           ),
//           SizedBox(width: 8),
//           Text(
//             duration ?? '',
//             style: TextStyle(fontSize: 12, color: isCurrentUser ? Colors.white : Colors.black),
//           ),
//         ],
//       ),
//     );
//   }
// }








// import 'package:flutter/material.dart';
// import 'package:google_fonts/google_fonts.dart';

// class MessageBubble extends StatelessWidget {
//   const MessageBubble({super.key});

//   @override
//   Widget build(BuildContext context) {
//     return Padding(
//       padding: EdgeInsets.all(10.0),
//       child: Row(
//         children: [
//           Flexible(
//             child: Column(
//               crossAxisAlignment: CrossAxisAlignment.end,
//               children: [
//                 Text(
//                   'Today',
//                   style: GoogleFonts.dmSans(
//                     textStyle: const TextStyle(
//                       fontWeight: FontWeight.w500,
//                       fontSize: 15,
//                       color: Color.fromRGBO(0, 14, 8, 1),
//                     ),
//                   ),
//                 ),
//                 const SizedBox(height: 5),
//                  Row(
//                   children: [
//                     SizedBox(width: 25),
//                     Flexible(
//                       child: Material(
//                         borderRadius: BorderRadius.only(
//                           topLeft: Radius.circular(30),
//                           bottomLeft: Radius.circular(30),
//                           bottomRight: Radius.circular(30),
//                         ),
//                         elevation: 5.0,
//                         color: Colors.blue,
//                         child: Padding(
//                           padding: const EdgeInsets.symmetric(
//                               horizontal: 20.0, vertical: 15.0),
//                           child: Text(
//                             'Hello lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua',
//                             style: TextStyle(
//                               color: Colors.white,
//                               fontSize: 15,
//                             ),
//                           ),
//                         ),
//                       ),
//                     ),
//                     SizedBox(width: 15)
//                   ],
//                 ),
//                 SizedBox(height: 5),
//                 Padding(
//                   padding: const EdgeInsets.only(left: 35.0),
//                   child: Row(
//                     children: [
//                       Text(
//                         '10:00 AM',
//                         style: GoogleFonts.dmSans(
//                           textStyle: const TextStyle(
//                             fontWeight: FontWeight.w500,
//                             fontSize: 12,
//                             color: Color.fromRGBO(0, 14, 8, 1),
//                           ),
//                         ),
//                       ),
//                     ],
//                   ),
//                 )
//               ],
//             ),
//           ),
//           SizedBox(width: 10),
//           Column(
//             children: [
//               Container(
//                 width: 40,
//                 height: 40,
//                 decoration: BoxDecoration(
//                   shape: BoxShape.circle,
//                   color: Colors.blue,
//                 ),
//                 child: Center(
//                   child: Image.asset('images/profile.png'),
//                 ),
//               ),
//               SizedBox(height: 140),
//             ],
//           )
//         ],
//       ),
//     );
//   }
// }
