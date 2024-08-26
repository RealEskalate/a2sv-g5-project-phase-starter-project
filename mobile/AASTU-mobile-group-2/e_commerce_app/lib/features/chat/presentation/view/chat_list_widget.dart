import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

Widget chatListItem(BuildContext context,{
  required ChatEntity chat,
  
}) {
  return ListTile(
    contentPadding: EdgeInsets.symmetric(vertical: 8.0, horizontal: 8.0),
    leading: CircleAvatar(
      radius: 28,
      backgroundImage: AssetImage('assets/smile.png'),
    ),
    title: Text(
      "title",
      style: GoogleFonts.lexend(
        textStyle: TextStyle(
          fontSize: 18,
          fontWeight: FontWeight.w500,
        ),
      ),
    ),
    subtitle: Text(
      "subtitle",
      style: GoogleFonts.lexend(
        textStyle: const TextStyle(
          fontSize: 12,
          fontWeight: FontWeight.w400,
          color: Color.fromARGB(255, 121, 124, 123),
        ),
      ),
    ),
    trailing: Column(
      crossAxisAlignment: CrossAxisAlignment.end,
      children: [
        Text(
          "time",
          style: GoogleFonts.lexend(
            textStyle: TextStyle(
              color: Color.fromARGB(255, 121, 124, 123),
              fontWeight: FontWeight.w400,
              fontSize: 12,
            ),
          ),
        ),
        if (2 > 0)
          Container(
            margin: EdgeInsets.only(top: 4.0),
            padding: EdgeInsets.all(6.0),
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              color: const Color.fromARGB(255, 63, 81, 243),
            ),
            child: Text(
              "2",
              style: GoogleFonts.abhayaLibre(
                textStyle: TextStyle(
                  fontWeight: FontWeight.w700,
                  fontSize: 12,
                  color: Colors.white,
                ),
              ),
            ),
          ),
      ],
    ),
    onTap: () {
      Navigator.pushNamed(context, '/chat',arguments: chat.id);
    },
  );
}
