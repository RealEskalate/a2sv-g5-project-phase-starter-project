import 'package:flutter/material.dart';

import '../widgets/chat_app_bar.dart';

class TextPage extends StatelessWidget {
  const TextPage({super.key});
  //here used the app bar widget instead of the appbar code 

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(),
      
      body: Column(
        children: [
          Expanded(
            child: ListView(
              padding: const EdgeInsets.all(8),
              children: const [
                // Incoming Message
                Text('Annei Ellison'),
                Text('Annei Ellison'),
                Text('You did your job well!'),
        
              ],
            ),
          ),
          
        ],
      ),
    );
  }

}