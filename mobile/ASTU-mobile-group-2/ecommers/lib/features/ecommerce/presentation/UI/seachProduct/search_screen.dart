


import 'package:flutter/material.dart';

import 'filter.dart';

import 'search_view.dart';

class SearchScreen extends StatefulWidget {
  const SearchScreen({super.key});


  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
  final searchInputController = TextEditingController();
  
  @override
  Widget build(BuildContext context) {
    
    return SafeArea(
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Search Product'),
          centerTitle: true,
          leading: GestureDetector(
            onTap: () => {
              Navigator.pop(context),
            },
            child: Container(
              width: 10,
              height: 10,
              margin: const EdgeInsets.only(left: 10),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(30),
                color: Colors.white
              ),
              child: const Icon(Icons.arrow_back_ios_new,color: Colors.blue,),
            ),
          ),
        ),
        body: Container(
          padding: const EdgeInsets.fromLTRB(20,15,20,10),
          child :  Column(
            children: [
              Row(
              children: [
                Expanded(
                  child: TextField(
                    controller: searchInputController,
                    onChanged: (value) => {
                      setState(() {
                        searchInputController.text = value;
                      })
                    },
                    decoration: InputDecoration(
                      hintText: 'Search',
                      enabledBorder: OutlineInputBorder(
                        borderSide: const BorderSide(
                          color: Colors.grey,
                        ),
                        borderRadius: BorderRadius.circular(10),
                      ),
                      suffixIcon: const Icon(Icons.arrow_forward,color:Colors.blue,size: 30,),
                    ),
                  ),
                ),
                GestureDetector(
                  onTap: () {
                    showModalBottomSheet(
                      context: context, 
                      builder: (BuildContext context){
                        return const Filter();
                      }
                      );
                  },
                  child: Container(
                    margin: const EdgeInsets.only(left: 10),
                    width: 50,
                    height: 50,
                    decoration: BoxDecoration(
                      color: Colors.blue,
                      borderRadius: BorderRadius.circular(10),
                    ),
                    child: const Icon(Icons.filter_list,size: 30,color:Colors.white,)),
                ),
              ],
            ),
              const SizedBox(height: 25,),
              Expanded(
                child: SearchView(text: searchInputController.text),
              ),
            ],
          )
        ),
      ),
    );
  }
}