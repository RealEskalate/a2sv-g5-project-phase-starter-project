import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';

class Bottomnavbar extends StatefulWidget {
  const Bottomnavbar({super.key});

  @override
  State<Bottomnavbar> createState() => _BottomnavbarState();
}

class _BottomnavbarState extends State<Bottomnavbar> {
  int _currentidx = 0;
  @override
  Widget build(BuildContext context) {
    return CurvedNavigationBar(
      height: MediaQuery.of(context).size.width * 0.17,
      buttonBackgroundColor: Color.fromRGBO(63, 81, 243, 1),
      color: Theme.of(context).colorScheme.onTertiary,
      backgroundColor: Colors.white,
      animationCurve: Curves.easeInOut,
      animationDuration: const Duration(milliseconds: 5000),
      // RGBO(63, 81, 243,1)
      onTap: (int newidx) {
        setState(() {
          _currentidx = newidx;
          Icon(
            Icons.home,
            color: Colors.white,
          );
        });
        switch (_currentidx) {
          case 0:
            // Navigator.pushNamed(context, '/home');
            break;
          case 1:
            // Navigator.pushNamed(context, '/add');
            break;
          case 2:
          // Navigator.pushNamed(context, '/search');
        }
      },

      items: [
        Icon(
          Icons.home,
          color: _currentidx == 0
              ? Colors.white
              : Theme.of(context).colorScheme.onSurface,
        ),
        // label: 'Home'
        // ),
        Icon(
          Icons.add,
          color: _currentidx == 1
              ? Colors.white
              : Theme.of(context).colorScheme.onSurface,
        ),
        // label: 'Add'
        // // ),
        // BottomNavigationBarItem(icon:
        Icon(
          Icons.search,
          color: _currentidx == 2
              ? Colors.white
              : Theme.of(context).colorScheme.onSurface,
        ),
        // label: 'Search'
        // ),
      ],
    );
  }
}
