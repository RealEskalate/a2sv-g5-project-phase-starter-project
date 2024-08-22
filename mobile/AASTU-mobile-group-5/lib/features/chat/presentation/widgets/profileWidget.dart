import 'package:flutter/material.dart';

class ProfileWidget extends StatelessWidget {
  const ProfileWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Container(
        alignment: Alignment.center,
        // width: 200,
        // height: 200,
        
        decoration:BoxDecoration(
          color: Colors.blue.shade300,
          shape: BoxShape.circle
        ),
        child: Card(
          shape: const CircleBorder(),                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  
          
          child: ClipOval(
            child: Container(
              child: SizedBox(
                width: 100,
                height: 100,
                 child: SizedBox(
                  width: 50,
                  height: 50,
                  child: Image.asset('assets/images/Alex.png',
                  fit: BoxFit.cover, 
                      width: double.infinity, 
                      height: double.infinity,
                  ),
                          
              ),
              ),
            ),
          ),
        ),
      ),
    );
    
  }
}