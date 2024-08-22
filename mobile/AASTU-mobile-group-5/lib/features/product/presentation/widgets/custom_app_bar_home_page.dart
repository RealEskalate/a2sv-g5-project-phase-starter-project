import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';
import '../../../../service_locator.dart';

import '../../../user/data/data_sources/user_local_data_source.dart';
import '../../../user/data/data_sources/user_name.dart';
import '../../../user/domain/use_case/logout_user.dart';
import 'dark_mode_toggle.dart';


class CustomAppBar extends StatefulWidget implements PreferredSizeWidget {
  const CustomAppBar({super.key});

  @override
  // ignore: library_private_types_in_public_api
  _CustomAppBarState createState() => _CustomAppBarState();

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}

class _CustomAppBarState extends State<CustomAppBar> {
  String userName = 'Name';

  @override
  void initState() {
    super.initState();
    _loadUserName();
  }

  Future<void> _loadUserName() async {
  try {
    
    final name = await fetchUserName();
    setState(() {
      userName = name;
    });
  } catch (e) {
      setState(() {
        userName = 'Guest'; // Fallback in case of an error
      });
  }
}
  @override
  Widget build(BuildContext context) {
    String formattedDate = DateFormat('yMMMMd').format(DateTime.now());

    final localDataSource = getIt<UserLocalDataSource>();
    final logOut = LogOut(localDataSource);
    return AppBar(
      automaticallyImplyLeading: false,
      title: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                // (204, 204, 204, 1)
                // Color(0xFF121212)
                color: Theme.of(context).brightness == Brightness.dark ? Colors.white : const Color.fromRGBO(204, 204, 204, 1),
              ),
              height: 50,
              width: 50,
            ),
            
            Expanded(
              child: Container(
                padding: const EdgeInsets.all(8),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Padding(
                      padding: const EdgeInsets.only(top: 0.0),
                      child: Text(
                        formattedDate,
                        style: GoogleFonts.syne(
                          fontSize: 13,
                          color: const Color.fromRGBO(170, 170, 170, 1),
                        ),
                      ),
                    ),
                    Row(
                      children: [
                        Text('Hello, ',
                            style: GoogleFonts.sora(
                              fontSize: 15,
                              color: const Color.fromRGBO(102, 102, 102, 1),
                              fontWeight: FontWeight.w500,
                            )),
                        Text(
                          userName,
                          style: TextStyle(
                            fontSize: 15,
                            color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
            const DarkThemeToggle(),
            Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                border: Border.all(
                  color: const Color.fromRGBO(217, 217, 217, 1),
                  width: 1.0,
                ),
              ),
              child: IconButton(
                icon: const Icon(Icons.logout_outlined,
                    color: Color.fromARGB(255, 164, 157, 157), size: 30),
                onPressed: () async {
                  await logOut();
                  // Optionally, navigate to the login screen or show a message
                  // ignore: use_build_context_synchronously
                  Navigator.pushReplacementNamed(context, '/signin');
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}
