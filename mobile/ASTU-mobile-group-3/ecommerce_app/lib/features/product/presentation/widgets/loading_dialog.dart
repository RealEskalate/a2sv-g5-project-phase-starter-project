import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';

import '../../../../core/themes/themes.dart';

class LoadingDialog extends StatelessWidget {
  const LoadingDialog({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.transparent,
      body: Center(
        child: Container(
          width: MediaQuery.of(context).size.width / 2,
          height: MediaQuery.of(context).size.width / 2,
          decoration: const BoxDecoration(
            color: MyTheme.ecWhite,
            borderRadius: BorderRadius.all(
              Radius.circular(50),
            ),
            boxShadow: [
              BoxShadow(color: MyTheme.ecBlack, spreadRadius: 2, blurRadius: 20)
            ],
          ),
          child: const Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              SpinKitWave(
                color: MyTheme.ecBlue,
                size: 50,
              ),
              SizedBox(
                height: 10,
              ),
              Text('Loading...'),
            ],
          ),
        ),
      ),
    );
  }
}
