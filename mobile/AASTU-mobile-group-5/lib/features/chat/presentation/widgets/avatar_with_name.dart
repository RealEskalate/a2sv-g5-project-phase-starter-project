import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'profile_widget.dart';

// ignore: camel_case_types
class AvatarWithName extends StatelessWidget {
  final List<String> names;
  const AvatarWithName({
    super.key,
    required this.names,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(
        right: 20,
      ),
      child: SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: Row(
          children: List.generate(names.length, (index) {
            return Padding(
              padding: const EdgeInsets.symmetric(horizontal: 8.0),
              child: Column(
                children: [
                  ProfileWidget(
                    isOnline: false,
                    iconUrl: 'assets/images/Alex.png',
                    isFirst: index == 0, // Set isFirst to true for the first element
                  ),
                  Text(
                    names[index],
                  style: GoogleFonts.poppins(
                      color: Colors.white,
                      fontSize: 12,
                    ),
                  ),
                ],
              ),
            );
          }),
        ),
      ),
    );
  }
}
