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
  void didChangeDependencies() {
    super.didChangeDependencies();
    // Synchronize current index based on current route
    String? currentRoute = ModalRoute.of(context)?.settings.name;
    if (currentRoute != null) {
      setState(() {
        _currentidx = _getIndexForRoute(currentRoute);
      });
    }
  }

  // Map routes to their respective index
  int _getIndexForRoute(String route) {
    switch (route) {
      case '/home':
        return 0;
      case '/add':
        return 1;
      case '/search':
        return 2;
      case '/HomeChat':
        return 3;
      default:
        return 0; // Default to home if route is unrecognized
    }
  }

  void _onTap(int newidx) {
    setState(() {
      _currentidx = newidx;
    });

    // Navigate to the corresponding page
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
      animationDuration: const Duration(milliseconds: 300),
      index: _currentidx,
      onTap: _onTap,
      items: [
        _buildIcon(Icons.home, 0),
        _buildIcon(Icons.add, 1),
        _buildIcon(Icons.search, 2),
        _buildIcon(Icons.textsms_outlined, 3),
      ],
    );
  }

  Widget _buildIcon(IconData icon, int index) {
    return Icon(
      icon,
      color: _currentidx == index ? Colors.white : Colors.black,
      size: _currentidx == index ? 32 : 24, // Increase size for active icon
    );
  }
}
