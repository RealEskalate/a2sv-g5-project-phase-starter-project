
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../../features/product/presentation/pages/home_page.dart';
import '../../screens/splash_screen.dart';

final routerDelegate = GoRouter(
  initialLocation: '/splash',
  routes: [
    GoRoute(
      path: '/splash',
      builder: (BuildContext context, GoRouterState state) => SplashScreen(),
    ),
    
    GoRoute(
      path: '/home',
      builder: (BuildContext context, GoRouterState state) => HomePage(),
    ),
  ],
);
