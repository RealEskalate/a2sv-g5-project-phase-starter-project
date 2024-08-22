import 'package:flutter/material.dart';
import 'package:ecommerce_app_ca_tdd/models/product.dart';
import 'resusetext.dart';
import 'package:google_fonts/google_fonts.dart';


class ItemCard extends StatelessWidget {
  final Product item;
  ItemCard({super.key, required this.item});


  @override
  Widget build(BuildContext context) {
    return Card(
      color: Colors.white,
      clipBehavior: Clip.antiAlias,
      child: GestureDetector(
        onTap: () {Navigator.pushNamed(context,'/detail', arguments: item);},
        child: Column(
          children: [
            AspectRatio(
              aspectRatio: 16 / 9,
              child: Image.asset(
                      item.imagePath,
                      fit: BoxFit.cover,
                    ),
            ),
            Container(
              color: Colors.white,
              padding: EdgeInsets.all(10),
              child: Column(
                children: [
                  Row(
                    children: [
                      Text(
                        item.name,
                        style: TextStyle(
                          fontSize: 20,
                          fontWeight: FontWeight.bold,
                          // fontWeight: FontWeight.bold
                        ),
                      ),
                      Spacer(),
                      Text(
                        "\$${item.price}",
                        style: TextStyle(
                          fontSize: 14,
                          fontWeight: FontWeight.bold,

                          // fontWeight: FontWeight.bold
                        ),
                      ),
                    ],
                  ),
                  Row(
                    children: [
                      Text(
                        "${item.category}",
                        style: TextStyle(
                          color: Color.fromARGB(255, 170, 170, 170),
                          fontSize: 12,
                          fontWeight: FontWeight.w400,
                        ),
                      ),
                      Spacer(),
                      Row(
                        children: [
                          Icon(
                            Icons.star,
                            color: Color.fromARGB(255, 255, 215, 0),
                            size: 20,
                          ),
                          Text(
                            "(${item.rating})",
                            style: TextStyle(
                              color: Color.fromARGB(255, 170, 170, 170),
                              fontSize: 12,
                              fontWeight: FontWeight.w400,
                            ),
                          )
                        ],
                      )
                    ],
                  ),
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}







// ignore: camel_case_types
class about_product extends StatefulWidget {
  const about_product({super.key});

  @override
  // ignore: library_private_types_in_public_api
  _CategoryState createState() => _CategoryState();
}

class _CategoryState extends State<about_product> {
  final TextEditingController _categoryController = TextEditingController();
    RangeValues _currentRangeValues = const RangeValues(0, 300);


  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.only(top: 38,right: 32,left: 32),
      child: Padding(
        padding: const EdgeInsets.all(2.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Container(
              
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  reusableTextpar("Category",FontWeight.w400,16),
                  SizedBox(
                    height: 40,
                    width: 366,
                    child: Container(
                      margin: EdgeInsets.only(top: 7),
                      child: TextField(
                        controller: _categoryController,
                        decoration: InputDecoration(
                          border: const OutlineInputBorder(),
                          
                          filled: true,
                          fillColor: Colors.white,
                        ),
                      ),
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 17),
            reusableTextpar("Price",FontWeight.w400,16), 
            const SizedBox(height: 10),
             Row(
                children: [
                  SizedBox(
                    height: 11,
                    width: 338,
                    child: RangeSlider(
                      values: _currentRangeValues,
                      min: 0,
                      max: 300,
                      divisions: 50,
                      labels: RangeLabels(
                        _currentRangeValues.start.round().toString(),
                        _currentRangeValues.end.round().toString(),
                      ),
                      
      
                      activeColor: Color.fromRGBO(63, 81, 243, 1), 
                      inactiveColor: Color.fromRGBO(217, 217, 217, 1),
                      onChanged: (RangeValues values) {
                        setState(() {
                          _currentRangeValues = values;
                        });
                      },
                    ),
                  ),
                ],
              ),
      
            const SizedBox(height: 65),
            Center(
              child: Container(
                // margin: EdgeInsets.only(right: 21),
                child: SizedBox(
                  width: 377,
                  height: 44,
                  child: ElevatedButton(
                    onPressed: () {
                    },
                    style: ElevatedButton.styleFrom(
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.all(Radius.circular(10))
                    ),
                    side: BorderSide(color: Color(0xff3F51F3)),
                    // overlayColor: Colors.red,
                    foregroundColor: Colors.white,
                    backgroundColor: Color(0xff3F51F3),
                                          ),
                    child: Text("APPLY",style: GoogleFonts.poppins(fontSize: 14,fontWeight: FontWeight.w500),),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}