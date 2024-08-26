import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';

class Bottomnavbar extends StatefulWidget {
  const Bottomnavbar({super.key});

  @override
  State<Bottomnavbar> createState() => _BottomnavbarState();
}

class _BottomnavbarState extends State<Bottomnavbar> {
  int _currentidx = 0;

  void _onTap(int newidx) {
    setState(() {
      _currentidx = newidx;
    });

    switch (_currentidx) {
      case 0:
        Navigator.pushNamed(context, '/home');
        break;
      case 1:
        Navigator.pushNamed(context, '/add');
        break;
      case 2:
        Navigator.pushNamed(context, '/search');
        break;
      case 3:
        Navigator.pushNamed(context, '/HomeChat');
        break;
    }
  }

  @override
  Widget build(BuildContext context) {
    return CurvedNavigationBar(
  height: MediaQuery.of(context).size.width * 0.17,
  buttonBackgroundColor: Color.fromRGBO(63, 81, 243, 1),
  color: const Color.fromARGB(255, 245, 244, 244),
  backgroundColor: Colors.white,
  animationCurve: Curves.easeInOut,
  animationDuration: const Duration(milliseconds: 300), // Adjusted to a standard animation duration
  index: _currentidx,
  onTap: _onTap,
  items: [
    Icon(
      Icons.home,
      color: _currentidx == 0 ? Colors.white : Colors.black,
    ),
    Icon(
      Icons.add,
      color: _currentidx == 1 ? Colors.white : Colors.black,
    ),
    Icon(
      Icons.search,
      color: _currentidx == 2 ? Colors.white : Colors.black,
    ),
     Icon(
      Icons.textsms_outlined,
      color: _currentidx == 3 ? Colors.white : Colors.black,
    ),
  ],
);
}
}
