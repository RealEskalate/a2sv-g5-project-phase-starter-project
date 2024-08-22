import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import 'home_page.dart';

class SearchPage extends StatefulWidget {
  const SearchPage({super.key});

  @override
  State<SearchPage> createState() => _SearchPageState();
}

class _SearchPageState extends State<SearchPage> {
  bool _showRangeSlider = false;
  RangeValues _currentRangeValues = const RangeValues(10, 100);

  final searchfieldController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return BlocListener<ProductBloc, ProductState>(
      listener: (context, state) {
        if (state is LoadedSingleProductState) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('single Product loaded successfully')),
          );
          // Navigator.pop(context); // Navigate back or to another screen
        } else if (state is ErrorState) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text(state.message)),
          );
        }
      },
      child: Scaffold(
        appBar: PreferredSize(
          preferredSize: const Size.fromHeight(70),
          child: AppBar(
            leading: IconButton(
              onPressed: () {
                Navigator.push(context,
                    MaterialPageRoute(builder: (context) => const HomePage()));
                context.read<ProductBloc>().add(LoadAllProductEvent());
              },
              icon: const Icon(
                Icons.arrow_back_ios,
                size: 18,
                color: Color.fromRGBO(63, 81, 243, 1),
              ),
            ),
            title: Padding(
              padding: const EdgeInsets.only(left: 50),
              child: Text(
                'Search product',
                style: GoogleFonts.poppins(
                  textStyle: const TextStyle(
                    fontWeight: FontWeight.w500,
                    color: Color.fromRGBO(62, 62, 62, 1),
                    fontSize: 16.0,
                  ),
                ),
              ),
            ),
          ),
        ),
        body: Column(
          children: [
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 32),
              child: Row(
                children: [
                  Expanded(
                    flex: 5,
                    child: TextFormField(
                      controller: searchfieldController,
                      decoration: InputDecoration(
                        labelText: 'Label',
                        labelStyle: GoogleFonts.poppins(
                          textStyle: const TextStyle(
                            fontSize: 14,
                            color: Color.fromRGBO(62, 62, 62, 1),
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                        suffixIcon: IconButton(
                          onPressed: () {
                            context.read<ProductBloc>().add(
                                GetSingleProductEvent(
                                    searchfieldController.text));
                          },
                          icon: const Icon(
                            Icons.arrow_forward,
                            color: Color.fromRGBO(63, 81, 243, 1),
                          ),
                        ),
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(12.0),
                          borderSide: const BorderSide(
                            color: Colors.grey,
                          ),
                        ),
                        enabledBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(12.0),
                          borderSide: const BorderSide(
                            color: Colors.grey,
                            width: 2.0,
                          ),
                        ),
                      ),
                      keyboardType: TextInputType.name,
                    ),
                  ),
                  const SizedBox(width: 5),
                  GestureDetector(
                    onTap: () {
                      setState(() {
                        _showRangeSlider = !_showRangeSlider;
                      });
                    },
                    child: Container(
                      height: 48,
                      width: 48,
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(13),
                        color: const Color.fromRGBO(63, 81, 243, 1),
                        shape: BoxShape.rectangle,
                        boxShadow: [
                          BoxShadow(
                            color: Colors.grey.withOpacity(0.5),
                            spreadRadius: 2,
                            blurRadius: 3,
                            offset: const Offset(0, 5),
                          ),
                        ],
                      ),
                      child: const Icon(
                        Icons.filter_list,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ],
              ),
            ),
            Expanded(
              // height: (MediaQuery.of(context).size.height) * 0.4,
              child: ListView(
                children: [
                  BlocBuilder<ProductBloc, ProductState>(
                    builder: (context, state) {
                      print(state);
                      if (state is LoadedSingleProductState) {
                        return Container(
                          margin: const EdgeInsets.only(
                              left: 32, right: 32, top: 16, bottom: 16),
                          height: 300,
                          decoration: BoxDecoration(
                            borderRadius: BorderRadius.circular(13),
                            color: Colors.white,
                            shape: BoxShape.rectangle,
                            boxShadow: [
                              BoxShadow(
                                color: Colors.grey.withOpacity(0.5),
                                spreadRadius: 3,
                                blurRadius: 5,
                                offset: const Offset(3, 5),
                              ),
                            ],
                          ),
                          child: Column(
                            mainAxisSize: MainAxisSize.max,
                            children: [
                              Image.network(state.product.imageUrl),
                              Container(
                                padding: const EdgeInsets.all(8),
                                child: Row(
                                  mainAxisAlignment:
                                      MainAxisAlignment.spaceBetween,
                                  children: [
                                    Text(
                                      state.product.name,
                                      style: GoogleFonts.poppins(
                                        textStyle: const TextStyle(
                                          fontWeight: FontWeight.w600,
                                          fontSize: 16,
                                          color: Color.fromRGBO(62, 62, 62, 1),
                                        ),
                                      ),
                                    ),
                                    Text(
                                      '\$${state.product.price}',
                                      style: GoogleFonts.poppins(
                                        textStyle: const TextStyle(
                                          fontWeight: FontWeight.w500,
                                          fontSize: 14.0,
                                        ),
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                              Padding(
                                padding:
                                    const EdgeInsets.symmetric(horizontal: 8.0),
                                child: Row(
                                  mainAxisAlignment:
                                      MainAxisAlignment.spaceBetween,
                                  children: [
                                    Text(
                                      'Men’s shoe',
                                      style: GoogleFonts.poppins(
                                        textStyle: const TextStyle(
                                          fontWeight: FontWeight.w600,
                                          color:
                                              Color.fromRGBO(170, 170, 170, 1),
                                          fontSize: 12.0,
                                          height: 0.5,
                                        ),
                                      ),
                                    ),
                                    Row(
                                      children: [
                                        const SizedBox(
                                          height: 24,
                                          width: 24,
                                          child: Center(
                                            child: Icon(
                                              Icons.star,
                                              size: 12,
                                              color: Color.fromRGBO(
                                                  255, 215, 0, 1),
                                            ),
                                          ),
                                        ),
                                        Text(
                                          '(4.0)',
                                          style: GoogleFonts.poppins(
                                            textStyle: const TextStyle(
                                              fontWeight: FontWeight.w400,
                                              fontSize: 12.0,
                                              color: Color.fromRGBO(
                                                  170, 170, 170, 1),
                                            ),
                                          ),
                                        ),
                                      ],
                                    ),
                                  ],
                                ),
                              ),
                            ],
                          ),
                        );
                      } else {
                        return Container();
                      }
                    },
                  )
                ],
              ),
            ),
            _showRangeSlider
                ? SizedBox(
                    child: Column(
                      children: [
                        _bottom_bar(context),
                        Container(
                          padding: const EdgeInsets.symmetric(horizontal: 32),
                          alignment: AlignmentDirectional.centerStart,
                          child: Text(
                            'Price',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w600,
                                fontSize: 16,
                                color: Color.fromRGBO(62, 62, 62, 1),
                              ),
                            ),
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.symmetric(horizontal: 32),
                          child: RangeSlider(
                            activeColor: const Color.fromRGBO(63, 81, 243, 1),
                            values: _currentRangeValues,
                            min: 0,
                            max: 100,
                            divisions: 100,
                            labels: RangeLabels(
                              _currentRangeValues.start.round().toString(),
                              _currentRangeValues.end.round().toString(),
                            ),
                            onChanged: (RangeValues values) {
                              setState(() {
                                _currentRangeValues = values;
                              });
                            },
                          ),
                        ),
                        const SizedBox(height: 20),
                        _apply_button(context),
                      ],
                    ),
                  )
                : Container(),

            // _item_list(context),
            // ListView(
            //   children: [
            //   ],
            // ),
            // _item_list(context),
            // _category(context),
          ],
        ),
      ),
    );
  }
}

Widget _apply_button(BuildContext context) {
  return Container(
    padding: EdgeInsets.symmetric(horizontal: 32),
    width: double.infinity,
    child: ElevatedButton(
      style: ElevatedButton.styleFrom(
        backgroundColor: const Color.fromRGBO(63, 81, 243, 1),
        elevation: 5,
        shadowColor: Colors.grey,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(12),
        ),
      ),
      onPressed: () {},
      child: Text(
        'Apply',
        style: GoogleFonts.poppins(
          color: Colors.white,
        ),
      ),
    ),
  );
}

Widget _bottom_bar(BuildContext context) {
  return Container(
    child: Column(
      children: [
        Divider(),
        Container(
          margin: const EdgeInsets.only(top: 20),
          padding: const EdgeInsets.symmetric(horizontal: 32),
          child: TextField(
            decoration: InputDecoration(
              labelText: 'Category',
              labelStyle: GoogleFonts.poppins(
                textStyle: const TextStyle(
                  fontSize: 14,
                  color: Color.fromRGBO(62, 62, 62, 1),
                  fontWeight: FontWeight.w500,
                ),
              ),
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(12.0),
                borderSide: const BorderSide(
                  color: Colors.grey,
                ),
              ),
              enabledBorder: OutlineInputBorder(
                borderRadius: BorderRadius.circular(12.0),
                borderSide: const BorderSide(
                  color: Colors.grey,
                  width: 2.0,
                ),
              ),
            ),
            keyboardType: TextInputType.name,
          ),
        ),
        const SizedBox(height: 20),
      ],
    ),
  );
}

// Widget _top_search_bar(BuildContext context) {
//   return Padding(
//     padding: const EdgeInsets.symmetric(horizontal: 32),
//     child: Row(
//       children: [
//         Expanded(
//           flex: 5,
//           child: TextField(
//             decoration: InputDecoration(
//               labelText: 'Label',
//               labelStyle: GoogleFonts.poppins(
//                 textStyle: const TextStyle(
//                   fontSize: 14,
//                   color: Color.fromRGBO(62, 62, 62, 1),
//                   fontWeight: FontWeight.w500,
//                 ),
//               ),
//               suffixIcon: IconButton(
//                 onPressed: () {
//                   context
//                       .read<ProductBloc>()
//                       .add(const GetSingleProductEvent());
//                 },
//                 icon: const Icon(
//                   Icons.arrow_forward,
//                   color: Color.fromRGBO(63, 81, 243, 1),
//                 ),
//               ),
//               border: OutlineInputBorder(
//                 borderRadius: BorderRadius.circular(12.0),
//                 borderSide: const BorderSide(
//                   color: Colors.grey,
//                 ),
//               ),
//               enabledBorder: OutlineInputBorder(
//                 borderRadius: BorderRadius.circular(12.0),
//                 borderSide: const BorderSide(
//                   color: Colors.grey,
//                   width: 2.0,
//                 ),
//               ),
//             ),
//             keyboardType: TextInputType.name,
//           ),
//         ),
//         const SizedBox(width: 5),
//         GestureDetector(
//           onTap: () {
//             // setState(() {
//             //     _showRangeSlider = !_showRangeSlider;
//             //   });
//           },
//           child: Container(
//             height: 48,
//             width: 48,
//             decoration: BoxDecoration(
//               borderRadius: BorderRadius.circular(13),
//               color: const Color.fromRGBO(63, 81, 243, 1),
//               shape: BoxShape.rectangle,
//               boxShadow: [
//                 BoxShadow(
//                   color: Colors.grey.withOpacity(0.5),
//                   spreadRadius: 2,
//                   blurRadius: 3,
//                   offset: const Offset(0, 5),
//                 ),
//               ],
//             ),
//             child: const Icon(
//               Icons.filter_list,
//               color: Colors.white,
//             ),
//           ),
//         ),
//       ],
//     ),
//   );
// }

// Widget _item_list(BuildContext context) {
//   return ListView(
//     children: [
//       const SizedBox(height: 10),
//       _itemList(context),
//       const SizedBox(height: 20),
//       _itemList(context),
//       const SizedBox(height: 20),
//       _itemList(context),
//       const SizedBox(height: 20),
//       _itemList(context),
//       const SizedBox(height: 20),
//     ],
//   );
// }

Widget _itemList(BuildContext context) {
  return GestureDetector(
    onTap: () {
      Navigator.pushNamed(context, '/details_page');
    },
    child: Container(
      margin: const EdgeInsets.only(left: 32, right: 32, top: 16, bottom: 16),
      height: 200,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(13),
        color: Colors.white,
        shape: BoxShape.rectangle,
        boxShadow: [
          BoxShadow(
            color: Colors.grey.withOpacity(0.5),
            spreadRadius: 3,
            blurRadius: 5,
            offset: const Offset(3, 5),
          ),
        ],
      ),
      child: Column(
        mainAxisSize: MainAxisSize.max,
        children: [
          Image.asset('images/shoes.png'),
          Container(
            padding: const EdgeInsets.all(8),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  'Derby Leather Shoes',
                  style: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                      fontWeight: FontWeight.w600,
                      fontSize: 16,
                      color: Color.fromRGBO(62, 62, 62, 1),
                    ),
                  ),
                ),
                Text(
                  '\$120',
                  style: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                      fontWeight: FontWeight.w500,
                      fontSize: 14.0,
                    ),
                  ),
                ),
              ],
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 8.0),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  'Men’s shoe',
                  style: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                      fontWeight: FontWeight.w600,
                      color: Color.fromRGBO(170, 170, 170, 1),
                      fontSize: 12.0,
                      height: 0.5,
                    ),
                  ),
                ),
                Row(
                  children: [
                    const SizedBox(
                      height: 24,
                      width: 24,
                      child: Center(
                        child: Icon(
                          Icons.star,
                          size: 12,
                          color: Color.fromRGBO(255, 215, 0, 1),
                        ),
                      ),
                    ),
                    Text(
                      '(4.0)',
                      style: GoogleFonts.poppins(
                        textStyle: const TextStyle(
                          fontWeight: FontWeight.w400,
                          fontSize: 12.0,
                          color: Color.fromRGBO(170, 170, 170, 1),
                        ),
                      ),
                    ),
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
