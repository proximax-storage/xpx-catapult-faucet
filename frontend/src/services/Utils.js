/**
 * Class for standard formats and validations
 */
export default class Utils {
  /**
   * Function to validate email
   * @param {*} email
   */
  static validEmail (email) {
    let re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    return re.test(email)
  };

  /**
  * Check if an address is valid
  *
  * @param {string} _address - An address
  *
  * @return {boolean} - True if address is valid, false otherwise
  */
 
  static isValid (_address) {

    if(_address == null)return false
    let address = _address.toString().toUpperCase().replace(/-/g, '');
    if (!address || address.length !== 40) {
        return false;
    }else{
        return true;
    }
  };

  /**
  * Add hyphens to a clean address
  *
  * @param {string} input - A PROXIMAX address
  *
  * @return {string} - A formatted PROXIMAX address
  */

  static addressForm (input) {
    return input && input.toUpperCase().replace(/-/g, '').match(/.{1,6}/g).join('-');
    };

  /**
  * Remove hyphens from an address
  *
  * @param {string} _address - An address
  *
  * @return {string} - A clean address
  */

  static clean (_address) {
    return _address.toUpperCase().replace(/-|\s/g,"");
  };

}
