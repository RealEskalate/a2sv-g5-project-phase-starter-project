import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
class BackgroundButton extends StatelessWidget {
  BackgroundButton({required this.title, super.key, this.callback});
  String title;

  final VoidCallback? callback;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: ElevatedButton(
        
        style: ButtonStyle(
          backgroundColor: MaterialStateProperty.all(Colors.blue),
          shape: MaterialStateProperty.all<RoundedRectangleBorder>(
              RoundedRectangleBorder(borderRadius: BorderRadius.circular(10))),
        ),
        onPressed: callback,
        child: Text(
          title,
          style: TextStyle(
              color: Colors.white, fontWeight: FontWeight.w600, fontSize: 14),
        ),
      ),
    );
  }
}





class TextFieldTitle extends StatelessWidget {
  TextFieldTitle(
      {super.key,
      this.title = '',
      
      this.pass = false,
      
      this.hint = null,

      required this.controller});

 
  String title;
  String? hint;

  bool pass;
  
  TextEditingController controller;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
     
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(height: 10,),
          title != ''
              ? Text(
                  title,
                  style: GoogleFonts.poppins(
                    color: Color.fromRGBO(111, 111, 111, 1),
                    fontSize: 16,
                    fontWeight: FontWeight.w400
                  )
                )
              : Container(),
          SizedBox(
            height: 10,
          ),
          Container(
            
            alignment: Alignment.center,
            decoration: BoxDecoration(
          
              borderRadius: BorderRadius.all(
                Radius.circular(10),
              ),
              color: Color.fromRGBO(245, 245, 245, 1),
            ),
            child: Center(
              child: TextField(
                obscureText: pass,
                
                controller: controller,
                decoration: InputDecoration(
                  border: InputBorder.none,
                  
                  hintText: hint,
                  hintStyle: GoogleFonts.poppins(
                    color: Color.fromRGBO(136, 136, 136, 1),
                    fontSize: 15,
                    fontWeight: FontWeight.w400
                  ),
                  contentPadding: EdgeInsets.only(left: 16, ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
class GoBack extends StatelessWidget {
  const GoBack({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return IconButton(
      onPressed: () {
        Navigator.pop(context);
      },
      icon: Icon(
        Icons.arrow_back_ios,
        size: 18,
        color: Color.fromARGB(255, 63, 81, 243),
      ),
    );
  }
}