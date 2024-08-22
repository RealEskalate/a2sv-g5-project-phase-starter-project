// ignore_for_file: prefer_const_constructors

import 'package:flutter/material.dart';
import 'package:get/get_navigation/src/snackbar/snackbar.dart';
import 'package:google_fonts/google_fonts.dart';

class CustomCard extends StatelessWidget {
  final String imagePath;
  final String name;
  final String description;
  final double price;
  

  const CustomCard({
    Key? key,
    required this.imagePath,
    required this.name,
    required this.description,
    required this.price,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () {
        Navigator.pushNamed(context, '/details');
      },
      child: Card(
        elevation: 2.0,
        margin: const EdgeInsets.all(8.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            
            ClipRRect(
                borderRadius: BorderRadius.only(
              topLeft: Radius.circular(16.0),
              topRight: Radius.circular(16.0)),
              child: Image(
              fit:BoxFit.cover,
                width: MediaQuery.of(context).size.width*1,
                height: MediaQuery.of(context).size.height*0.3,
                image: AssetImage(imagePath)),
            ), // For network image
            SizedBox(height: 5.0),
            Padding(
              padding: const EdgeInsets.only(left:16.0,top: 4,bottom: 12,right: 21),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [Text(name,style: GoogleFonts.poppins(fontSize: 20.0,fontWeight: FontWeight.w500,),),
                          SizedBox(height: 8.0),
                  Text(description,style: GoogleFonts.poppins(fontSize: 12.0,fontWeight: FontWeight.w400,color: Color.fromRGBO(217, 217, 217, 1))),
                    ],
                  ),
                SizedBox(height: 8.0),
                  Column(
                    children: [
                      Text(
                        '\$'+price.toString(),
                        style: GoogleFonts.sora(fontSize: 14,fontWeight: FontWeight.w500),
                      ),
                      Row(children: [
                        Icon(Icons.star,size: 24 ,color: Color.fromARGB(255, 255, 215, 0),),
                        Text("(4.0)",style:GoogleFonts.sora(fontSize: 12,fontWeight: FontWeight.w400,color: Color.fromARGB(255, 170, 170, 170))),
                        
                      ],)
                    ],
                  ),
                ],
              
              
              
              
              ),
            ),
          ],
        ),
      ),
    );
  }
}


final List<CustomCard> items = [
  CustomCard(
    imagePath: 'assets/mike.jpg',
    name: 'Travis x Nike Lows',
    description: 'Men' + "'"+ 's'+' shoes',
    price: 29.99,

  ),
  CustomCard(
    imagePath: 'assets/jeff.jpg',
    name: 'Nike Mids',
    description:'Men' + "'"+ 's'+' shoes',
    price: 39.99,
  ),CustomCard(
    imagePath: 'assets/zachary-keimig-V_aqnVipk-s-unsplash.jpg',
    name: 'Nike Lows',
    description:'Men ' + "'"+ 's'+' shoes',
    price: 39.99,
  ),CustomCard(
    imagePath: 'assets/hightpop.jpg',
    name: 'Nike High Tops',
    description:'Men ' + "'"+ 's'+' shoes',
    price: 59.99,
  ),
  // Add more items as needed
];






















// import "package:flutter/material.dart";
// import 'package:google_fonts/google_fonts.dart';

// class ShoeCard extends StatelessWidget {
//   const ShoeCard({super.key});
  

//   @override
//   Widget build(BuildContext context) {
//     return  Card(
//     child: Column(
//       children: [
//         ClipRRect(
//             borderRadius: BorderRadius.only(
//                 topLeft: Radius.circular(20.0),
//                 topRight: Radius.circular(20.0)),
//                 child:Image(
//                           // height: 300.0,
//                           image: AssetImage('assets/shoes.jpg'),
                        
//                         ),
//         ),
        
//         Row(
//           mainAxisAlignment: MainAxisAlignment.spaceBetween,
//             children: [
//               Padding(
//                 padding: const EdgeInsets.all(11.0),
//                 child: Column(
//                   crossAxisAlignment: CrossAxisAlignment.start,
//                   children: [
//                     Padding(
//                       padding: const EdgeInsets.only(bottom: 10.0),
//                       child: Text("Derby Leather Shoes",style:TextStyle(fontFamily: "Poppins", fontSize: 20,fontWeight: FontWeight.w500),),
//                     ),
//                     Text("Men" +"’" +"s"  + " shoe",style: TextStyle(fontFamily: "Poppins", fontSize: 12,fontWeight: FontWeight.w400,color: Color.fromARGB(255, 170, 170, 170)),),
//                   ],
//                 ),
//               ),
//               Padding(
//                 padding: const EdgeInsets.all(11.0),
//                 child: Column(
//                   crossAxisAlignment: CrossAxisAlignment.end,
//                   children: [
//                     Padding(
//                       padding: const EdgeInsets.only(bottom: 10.0),
//                       child: Text("\$120",style: GoogleFonts.sora(fontSize: 14,fontWeight: FontWeight.w500),),
//                     ),
//                     Row(children: [
//                       Icon(Icons.star,size: 24 ,color: Color.fromARGB(255, 255, 215, 0),),
//                       Text("(4.0)",style:GoogleFonts.sora(fontSize: 12,fontWeight: FontWeight.w400,color: Color.fromARGB(255, 170, 170, 170))),
                      
//                     ],)
//                   ],
//                 ),
//               )
//             ],
//         )
//       ]
//       ,),
//   );
// }
// }



// // Custom widget

// Card customKicks(Image photo,String name,String attire,String price,Icon star,String rating){
//   return Card(
//     child: Column(
//       children: [
//         ClipRRect(
//             borderRadius: BorderRadius.only(
//                 topLeft: Radius.circular(20.0),
//                 topRight: Radius.circular(20.0)),
//                 child:Image(
//                           // height: 300.0,
//                           image: AssetImage(photo),
                        
//                         ),
//         ),
        
//         Row(
//           mainAxisAlignment: MainAxisAlignment.spaceBetween,
//             children: [
//               Padding(
//                 padding: const EdgeInsets.all(11.0),
//                 child: Column(
//                   crossAxisAlignment: CrossAxisAlignment.start,
//                   children: [
//                     Padding(
//                       padding: const EdgeInsets.only(bottom: 10.0),
//                       child: Text(,style:TextStyle(fontFamily: "Poppins", fontSize: 20,fontWeight: FontWeight.w500),),
//                     ),
//                     Text("Men" +"’" +"s"  + " shoe",style: TextStyle(fontFamily: "Poppins", fontSize: 12,fontWeight: FontWeight.w400,color: Color.fromARGB(255, 170, 170, 170)),),
//                   ],
//                 ),
//               ),
//               Padding(
//                 padding: const EdgeInsets.all(11.0),
//                 child: Column(
//                   crossAxisAlignment: CrossAxisAlignment.end,
//                   children: [
//                     Padding(
//                       padding: const EdgeInsets.only(bottom: 10.0),
//                       child: Text("\$120",style: GoogleFonts.sora(fontSize: 14,fontWeight: FontWeight.w500),),
//                     ),
//                     Row(children: [
//                       Icon(Icons.star,size: 24 ,color: Color.fromARGB(255, 255, 215, 0),),
//                       Text("(4.0)",style:GoogleFonts.sora(fontSize: 12,fontWeight: FontWeight.w400,color: Color.fromARGB(255, 170, 170, 170))),
                      
//                     ],)
//                   ],
//                 ),
//               )
//             ],
//         )
//       ]
//       ,),
//   )
// }

  