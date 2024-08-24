// import 'dart:ui';
// import 'package:flutter/material.dart';
// import 'package:intl/intl.dart';  // Add this import for formatting time

// import 'user_avater.dart';

// class Chat extends StatefulWidget {
//   final String message;
//   const Chat({super.key, required this.message});

//   @override
//   State<Chat> createState() => _ChatState();
// }

// class _ChatState extends State<Chat> {
//     bool isSelf = false;
//     DateTime sendTime = DateTime.now();

//   @override
//     String formattedTime = DateFormat('hh:mm a').format(sendTime); // Format time with AM/PM

//   Widget build(BuildContext context) {
//     return Column(children: [
//     SizedBox(height: 30),
//     isSelf ? 
//     Positioned(
//       right: 0,
//       child: Row(
//         crossAxisAlignment: CrossAxisAlignment.start,
//         mainAxisAlignment: MainAxisAlignment.end,
//         children: [
//           Text('You', style: TextStyle(color: Color.fromARGB(255, 11, 11, 11), fontSize: 14, fontFamily: 'Poppins', fontWeight: FontWeight.w900)),
//           SizedBox(width: 5),
//           UserAvater(image:'assets/images/avater.png', online: false, story: false)
//         ],
//       ),
//     )
//     :
//     Positioned(
//       left: 0,
//       child: Row(
//         crossAxisAlignment: CrossAxisAlignment.start,
//         mainAxisAlignment: MainAxisAlignment.start,
//         children: [
//           UserAvater(image:'assets/images/avater.png', online: false, story: false),
//           SizedBox(width: 5),
//           Text('Sabila Sayma', style: TextStyle(color: Color.fromARGB(255, 11, 11, 11), fontSize: 14, fontFamily: 'Poppins', fontWeight: FontWeight.w900)),
//         ],
//       ),
//     ),
//     SizedBox(height: 4),
//     ClipRRect(
//       borderRadius: BorderRadius.only(
//         topRight: isSelf ? Radius.circular(0) : Radius.circular(18),
//         topLeft: isSelf ? Radius.circular(18) : Radius.circular(0),
//         bottomLeft: Radius.circular(18),
//         bottomRight: Radius.circular(18)
//       ),


//       child: Container(
//         width: 280,
//         child: Column(
//           children: [
//           decoration: BoxDecoration(
//             color: isSelf ? Color(0xFF3E50F3): Color(0xFFF2F7FB),
//           ),
//           child: Padding(
//             padding: const EdgeInsets.symmetric(horizontal: 14, vertical:  14.0),
//             child: Text(
//               widget.message,
//               style: TextStyle(
//                 color: isSelf ? Colors.white : Color.fromARGB(255, 36, 36, 36),
//                 fontSize: 12,
//                 fontFamily: 'Poppins',
//                 fontWeight: FontWeight.w400
//               ),
//             ),
//           ),
//     child: Text('${sendTime.hour}:${sendTime.minute} ', 
//     style: TextStyle(
//       color: Color.fromARGB(255, 11, 11, 11), 
//       fontSize: 10, 
//       fontFamily: 'Poppins', 
//       fontWeight: FontWeight.w400),
//     textAlign: TextAlign.end,
//     ),
//     ],),
//       ),
//     ),
//   ],);
//   }
// }

import 'package:flutter/material.dart';
import 'user_avater.dart';
import 'package:intl/intl.dart';  // Add this import for formatting time

class Chat extends StatefulWidget {
  final String message;
  final bool isSelf;
  const Chat({Key? key, required this.message, this.isSelf = true}) : super(key: key);

  @override
  _ChatState createState() => _ChatState();
}

class _ChatState extends State<Chat> {
  DateTime sendTime = DateTime.now();

  @override
  Widget build(BuildContext context) {
    String formattedTime = DateFormat('hh:mm a').format(sendTime); // Format time with AM/PM

    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 10.0),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisAlignment: widget.isSelf ? MainAxisAlignment.end : MainAxisAlignment.start,
        children: [
          if (!widget.isSelf) UserAvater(image: 'assets/images/avater.png', online: false, story: false),
          SizedBox(width: 5),
          Column(
            crossAxisAlignment: widget.isSelf ? CrossAxisAlignment.end : CrossAxisAlignment.start,
            children: [
              Text(
                widget.isSelf ? 'You' : 'Sabila Sayma',
                style: TextStyle(
                  color: Color.fromARGB(255, 11, 11, 11),
                  fontSize: 14,
                  fontFamily: 'Poppins',
                  fontWeight: FontWeight.w900,
                ),
              ),
              SizedBox(height: 4),
              ClipRRect(
                borderRadius: BorderRadius.only(
                  topRight: widget.isSelf ? Radius.circular(0) : Radius.circular(18),
                  topLeft: widget.isSelf ? Radius.circular(18) : Radius.circular(0),
                  bottomLeft: Radius.circular(18),
                  bottomRight: Radius.circular(18),
                ),
                child: Container(
                  width: 280,
                  decoration: BoxDecoration(
                    color: widget.isSelf ? Color(0xFF3E50F3) : Color(0xFFF2F7FB),
                  ),
                  child: Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 14, vertical: 14.0),
                    child: Text(
                      widget.message,
                      style: TextStyle(
                        color: widget.isSelf ? Colors.white : Color.fromARGB(255, 36, 36, 36),
                        fontSize: 12,
                        fontFamily: 'Poppins',
                        fontWeight: FontWeight.w400,
                      ),
                    ),
                  ),
                ),
              ),
              SizedBox(height: 4),
              Text(
                formattedTime,
                style: TextStyle(
                  color: Color.fromARGB(255, 11, 11, 11),
                  fontSize: 10,
                  fontFamily: 'Poppins',
                  fontWeight: FontWeight.w400,
                ),
              ),
            ],
          ),
          if (widget.isSelf) SizedBox(width: 5),
          if (widget.isSelf) UserAvater(image: 'assets/images/avater.png', online: false, story: false),
        ],
      ),
    );
  }
}
