import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'profile_widget.dart';

// ignore: camel_case_types
class UserConversationWidget extends StatelessWidget {
  final String name;
  const UserConversationWidget({
    super.key,
    required this.name,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(vertical: 8), // Adjust vertical margin
      padding: const EdgeInsets.symmetric(
          horizontal: 8), // Adjust horizontal padding
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            children: [
              const ProfileWidget(
                isOnline: true,
                iconUrl: 'assets/images/Alex.png',
                isInAppBar: false,
                curStatus: false,
                isFirst: false,
              ),
              const SizedBox(
                width: 15,
              ),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisSize:
                    MainAxisSize.min, // Minimize the height of the column
                children: [
                  Text(
                    name,
                    style: GoogleFonts.poppins(
                      color: Colors.black,
                      fontSize: 18,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                  Text(
                    'How are you today?',
                    style: GoogleFonts.poppins(
                      color: Color.fromRGBO(121, 124, 123, 1),
                      fontSize: 12,
                      fontWeight: FontWeight.w400,
                    ),
                  ),
                ],
              ),
            ],
          ),
          Column(
            mainAxisSize: MainAxisSize.min, // Minimize the height of the column
            children: [
              Text(
                '2 mins ago',
                style: GoogleFonts.poppins(
                  color: Color.fromRGBO(121, 124, 123, 1),
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ),
              SizedBox(
                height: 8,
              ),
              CircleAvatar(
                radius: 15,
                backgroundColor: Color.fromRGBO(73, 140, 240, 1),
                child: Text(
                  '4',
                  style: GoogleFonts.abhayaLibre(
                    fontWeight: FontWeight.w700,
                    fontSize: 18,
                    height: 1, // Line height
                    color: Color.fromRGBO(255, 255, 255, 1),
                  ),
                  textAlign: TextAlign.center,
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
