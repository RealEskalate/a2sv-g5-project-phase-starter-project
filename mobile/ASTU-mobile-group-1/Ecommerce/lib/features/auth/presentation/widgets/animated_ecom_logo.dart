import 'package:flutter/material.dart';

import 'ecom_logo.dart';

class AnimatedEcomLogo extends StatefulWidget {
  const AnimatedEcomLogo({super.key});

  @override
  State<AnimatedEcomLogo> createState() => _AnimatedEcomLogoState();
}

class _AnimatedEcomLogoState extends State<AnimatedEcomLogo>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  late Animation<Offset> _offsetAnimation;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(seconds: 1),
    );
    _offsetAnimation = Tween<Offset>(
      begin: const Offset(2, 0),
      end: Offset.zero,
    ).animate(CurvedAnimation(
      parent: _controller,
      curve: Curves.easeInOut,
    ));
    _controller.forward();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SlideTransition(
      position: _offsetAnimation,
      child: const EcomLogo(
        fontSize: 100,
        width: 250,
        backgroundColor: Colors.white,
        boarderRadius: 30,
      ),
    );
  }
}
