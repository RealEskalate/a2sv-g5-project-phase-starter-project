import "package:flutter/material.dart";
import "package:google_fonts/google_fonts.dart";
import "package:shared_preferences/shared_preferences.dart";

class LogoutScreen extends StatelessWidget {
  const LogoutScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return  Scaffold(
      body: Center(
        child: ElevatedButton(
          style: ElevatedButton.styleFrom(
            backgroundColor: Colors.red,
            textStyle: TextStyle(fontSize: 20, fontFamily: 'Montserrat'),
            foregroundColor: Colors.white,
          ),
          onPressed: () {
            showDialog(
              context: context, 
              builder: (context)=> AlertDialog(
                title: Text("Are you sure you want to logout ?",style: GoogleFonts.poppins(fontSize: 15),),
                actions: [
                  TextButton(onPressed: (){
                    Navigator.pop(context);
                  }, child: Text("Cancel")),
                  TextButton(
                    onPressed: (){
                      logOut();
                      Navigator.pushNamedAndRemoveUntil(context, '/login', (route) => true);
                    },
                    child: Text("Log-Out")
                    )
                ],
                                          )
                                        );
          },
          child: Text("Log-Out"),)
      ),
    );
  }
}

Future<void> logOut() async {
  SharedPreferences remove = await SharedPreferences.getInstance();
  remove.remove('access_token');
}
Future<bool?> checkToken() async {
  SharedPreferences remove = await SharedPreferences.getInstance();
  var result = remove.getString('access_token');
  if (result != null){
    return false;
  }else{
  return true;

  }
}