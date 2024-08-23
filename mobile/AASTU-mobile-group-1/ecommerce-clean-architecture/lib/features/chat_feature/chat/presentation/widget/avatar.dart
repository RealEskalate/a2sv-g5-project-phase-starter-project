import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class avatar extends StatelessWidget{
  String name;
  avatar({super.key,required this.name});
  @override
  Widget build(BuildContext context) {
    return Center(
          child: Column(
            children: [
              Stack(
                alignment: Alignment.center,
                children: [
                  Container(
                    width: 70, // Size of the circular ring
                    height: 70, // Size of the circular ring
                    decoration: BoxDecoration(
                      shape: BoxShape.circle,
                      border: Border.all(
                        color: Colors.white, // Border color
                        width: 1, // Border width
                      ),
                    ),
                  ),
                  CircleAvatar(
                    radius: 30, // Radius of the avatar
                    backgroundImage: AssetImage('assets/images/avatar.jpg'),
                  ),
                ],
              ),
              SizedBox(height: 7),
              Text(
                name,
                style: GoogleFonts.poppins(
                  textStyle: TextStyle(
                    color: Colors.white,
                    fontSize: 15,
                    fontWeight: FontWeight.w400,
                  ),
                ),
              ),
            ],
          ),
        );
  }
}
