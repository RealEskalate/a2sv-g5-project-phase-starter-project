import 'package:flutter/material.dart';

class ProfileWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Container(
        alignment: Alignment.center,
        
        
        decoration:BoxDecoration(
          color: Colors.blue.shade300,
          shape: BoxShape.circle
        ),
        child: Card(
          shape: CircleBorder(),                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  
          
          child: ClipOval(
            child: Container(
              child: Container(
                width: 100,
                height: 100,
                 child: Container(
                  width: 600,
                  height: 600,
                  child: Image.asset('assets/images/Alex.png',
                  fit: BoxFit.fill, 
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