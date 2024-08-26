import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';

class Bottomnavbar extends StatefulWidget {
  const Bottomnavbar({super.key});

  @override
  State<Bottomnavbar> createState() => _BottomnavbarState();
}

class _BottomnavbarState extends State<Bottomnavbar> {
  int _currentIdx = 0;

  @override
  Widget build(BuildContext context) {
    return CurvedNavigationBar(
      height: MediaQuery.of(context).size.width * 0.17,
      buttonBackgroundColor: const Color.fromRGBO(63, 81, 243, 1),
      color: const Color.fromARGB(255, 245, 244, 244),
      backgroundColor: Colors.white,
      animationCurve: Curves.easeInOut,
      animationDuration: const Duration(milliseconds: 300),
      onTap: (int newIdx) {
        setState(() {
          _currentIdx = newIdx;
        });

        // Navigate based on the selected index
        switch (_currentIdx) {
          case 0:
            Navigator.pushNamed(context, '/home');
            break;
          case 1:
            Navigator.pushNamed(context, '/add');
            break;
          case 2:
            Navigator.pushNamed(context, '/search');
            break;
        }
      },
      items: [
        Icon(
          Icons.home,
          color: _currentIdx == 0 ? Colors.white : Colors.black,
        ),
        Icon(
          Icons.add,
          color: _currentIdx == 1 ? Colors.white : Colors.black,
        ),
        Icon(
          Icons.search,
          color: _currentIdx == 2 ? Colors.white : Colors.black,
        ),
      ],
    );
  }
}
